package api

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Create(ctx context.Context, req *activitycomm.CreateActivityRequest) (*activitycomm.CreateActivityReply, error) {
	// parse inputs
	if err := s.validateActivityType(ctx, req.GetActTypeID()); err != nil {
		return nil, err
	}

	period, err := s.validatePeriod(req.GetPeriod())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error creating activity: %v", err)
	}

	// period := req.GetPeriod()
	// create store data
	act := storage.Activity{
		Period:         *period,
		ActivityTypeID: req.GetActTypeID(),
		OrderID:        req.GetOrderID(),
	}

	// persista data
	id, err := s.repo.Create(ctx, act)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating activity: %v", err)
	}

	// reply
	return &activitycomm.CreateActivityReply{
		Id: id.String(),
	}, nil
}

func (s *activityServer) validatePeriod(period *activitycomm.Interval) (*storage.Interval, error) {
	if period == nil {
		return nil, fmt.Errorf("Period is not defined")
	}
	if period.From == nil {
		return nil, fmt.Errorf("Period's From is not defined")
	}
	if period.To == nil {
		return nil, fmt.Errorf("Period's To is not defined")
	}

	from, err := ptypes.Timestamp(period.From)
	if err != nil {
		return nil, fmt.Errorf("Period's From is invalid: %v", period.From)
	}

	to, err := ptypes.Timestamp(period.To)
	if err != nil {
		return nil, fmt.Errorf("Period's To is invalid: %v", period.To)
	}

	return &storage.Interval{
		From: from,
		To:   to,
	}, nil
}

func (s *activityServer) validateActivityType(ctx context.Context, actTypeIDStr string) error {
	_, err := uuid.Parse(actTypeIDStr)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "error creating activity, ActivityType ID is invalid: %v", err)
	}

	// validate inputs
	rpl, err := s.actTypeCli.Exist(ctx, &activitytypecomm.ExistActivityTypeRequest{Id: actTypeIDStr})
	if err != nil {
		return status.Errorf(codes.Unavailable, "error contacting ActivityType service: %v", err)
	}
	if !rpl.Exists {
		return status.Errorf(codes.NotFound, "ActivityType with id %v do not exists", actTypeIDStr)
	}

	return nil
}
