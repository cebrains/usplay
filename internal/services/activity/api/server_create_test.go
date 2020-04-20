package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_CreateHappyPath(t *testing.T) {
	// arrange
	activityID := uuid.New()
	activity := storage.Activity{
		ActivityTypeID: uuid.New().String(),
		OrderID:        uuid.New().String(),
		Code:           "Activity Code",
		Description:    "Activity Description",
		Name:           "Activity Name",
	}
	store := &activityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	reply, err := svr.Create(ctx, &activitycomm.CreateActivityRequest{
		ActTypeID:   activity.ActivityTypeID,
		OrderID:     activity.OrderID,
		Code:        activity.Code,
		Description: activity.Description,
		Name:        activity.Name,
	})

	// assert
	if err != nil {
		t.Fatalf("error invoking create: %v", err)
	}

	if expected, provided := store.CreateResult.ID.String(), reply.GetId(); expected != provided {
		t.Errorf("expected id %s, provided %s", expected, provided)
	}
}

func Test_CreateInvalidActivityTypeID(t *testing.T) {
	// arrange
	activityID, orderID := uuid.New(), uuid.New()
	activity := storage.Activity{
		ID:          activityID.String(),
		OrderID:     orderID.String(),
		Code:        "Activity Code",
		Description: "Activity Description",
		Name:        "Activity Name",
	}
	store := &activityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: true,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	_, err := svr.Create(ctx, &activitycomm.CreateActivityRequest{
		ActTypeID:   "",
		Code:        activity.Code,
		Description: activity.Description,
		Name:        activity.Name,
	})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking create with no id is not provided")
	}

	statusErr := status.Convert(err)
	if statusErr == nil {
		t.Fatalf("provided error is not a status.Status error: %v", err)
	}

	if statusErr.Code() != codes.InvalidArgument {
		t.Errorf("provided error do not present the InvalidArgument code as expected, but instead presents %s", statusErr.Code().String())
	}
}

func Test_CreateNotExistingActivityTypeID(t *testing.T) {
	// arrange
	activityID := uuid.New()
	activity := storage.Activity{
		ActivityTypeID: uuid.New().String(),
		OrderID:        uuid.New().String(),
		Code:           "Activity Code",
		Description:    "Activity Description",
		Name:           "Activity Name",
	}
	store := &activityTestRepo{
		CreateResult: struct {
			ID  *uuid.UUID
			Err error
		}{
			ID:  &activityID,
			Err: nil,
		},
	}
	svr := api.NewActivityServer(
		store,
		&actTestClient{
			WaitTime: time.Duration(0),
			ExistResult: struct {
				Err   error
				Reply activitytypecomm.ExistActivityTypeReply
			}{
				Err: nil,
				Reply: activitytypecomm.ExistActivityTypeReply{
					Exists: false,
				},
			},
		},
		&orderTestClient{
			ExistResult: struct {
				Err   error
				Reply ordercomm.ExistOrderReply
			}{
				Err: nil,
				Reply: ordercomm.ExistOrderReply{
					Exists: true,
				},
			},
		},
		1*time.Second,
	)
	ctx := context.Background()

	// act
	_, err := svr.Create(ctx, &activitycomm.CreateActivityRequest{
		ActTypeID:   activity.ActivityTypeID,
		Code:        activity.Code,
		Description: activity.Description,
		Name:        activity.Name,
	})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking create with non-existing activitytype id provided")
	}

	statusErr := status.Convert(err)
	if statusErr == nil {
		t.Fatalf("provided error is not a status.Status error: %v", err)
	}

	if statusErr.Code() != codes.NotFound {
		t.Errorf("provided error do not present the NotFound code as expected, but instead presents %s", statusErr.Code().String())
	}
}
