package googlesearch

import (
	"context"

	"github.com/RafaZeero/brand_monitor/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Collection
}

func NewStore(db *mongo.Collection) *Store {
	return &Store{db: db}
}

func (s *Store) CreateSearch(
	ctx context.Context,
	payload *types.CreateSearchPayload,
) error {
	_, err := s.db.InsertOne(context.Background(), payload)
	return err
}

func (s *Store) GetSearches(
	context.Context,
) ([]*types.Search, error) {
	return nil, nil
}

func (s *Store) GetSearchByID(
	context.Context,
	primitive.ObjectID,
) (*types.Search, error) {
	return nil, nil
}
