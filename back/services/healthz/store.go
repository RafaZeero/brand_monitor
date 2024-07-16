package healthz

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) Ping() error {
	return s.db.Ping(context.TODO(), nil)
}
