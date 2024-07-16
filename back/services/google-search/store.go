package googlesearch

import (
	"context"
	"time"

	"github.com/RafaZeero/brand_monitor/types"
	"go.mongodb.org/mongo-driver/bson"
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
	_, err := s.db.InsertOne(context.Background(), &types.Search{
		ID:          primitive.NewObjectID(),
		Term:        payload.Term,
		Competitors: payload.Competitors,
		CreatedAt:   time.Now(),
	})
	return err
}

func (s *Store) GetSearches(
	ctx context.Context,
) ([]*types.Search, error) {
	cur, err := s.db.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	var searches []*types.Search
	if err = cur.All(ctx, &searches); err != nil {
		return nil, err
	}

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	cur.Close(ctx)

	if len(searches) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return searches, nil
}

func (s *Store) GetSearchByID(
	context.Context,
	primitive.ObjectID,
) (*types.Search, error) {
	return nil, nil
}
