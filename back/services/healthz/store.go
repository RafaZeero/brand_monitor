package healthz

import (
	"context"

	"github.com/RafaZeero/brand_monitor/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Collection
}

func NewStore(db *mongo.Collection) *Store {
	return &Store{db: db}
}

func (s *Store) AddData(ctx context.Context, data *types.TestAddData) error {
	_, err := s.db.InsertOne(ctx, data)
	return err
}
