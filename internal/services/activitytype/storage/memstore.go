package storage

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/usplay/pkg/repoerrors"
	"github.com/google/uuid"
)

type memoryStore struct {
	data map[uuid.UUID]ActivityType
}

// NewInMemoryStore in memory repository for activities
func NewInMemoryStore() Repository {
	return &memoryStore{
		data: map[uuid.UUID]ActivityType{},
	}
}

// Create store a new ActivityType
func (s *memoryStore) Create(ctx context.Context, a ActivityType) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	a.ID = id
	s.data[id] = a
	return &id, nil
}

// Exist checks if an activitytype exists
func (s *memoryStore) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	_, ok := s.data[id]
	return &ok, nil
}

// Read get an ActivityType by id
func (s *memoryStore) Read(ctx context.Context, id uuid.UUID) (*ActivityType, error) {
	act, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("ActivityType with id %v not found", id)
	}
	return &act, nil
}

// Update updates an ActivityType in the store if present
func (s *memoryStore) Update(ctx context.Context, a ActivityType) error {
	if _, ok := s.data[a.ID]; !ok {
		return &repoerrors.NotFoundError{ID: a.ID}
	}

	s.data[a.ID] = a
	return nil
}

// Delete removes an ActivityType from the store if present
func (s *memoryStore) Delete(ctx context.Context, id uuid.UUID) error {
	delete(s.data, id)
	return nil
}

// List returns all the activities in the store
func (s *memoryStore) List(ctx context.Context, ids []uuid.UUID) (Activities, error) {
	list := make([]ActivityType, len(ids), len(s.data))
	for i, v := range ids {
		a, ok := s.data[v]
		if !ok {
			return nil, fmt.Errorf("ID %v not found", v)
		}
		list[i] = a
	}
	return list, nil
}
