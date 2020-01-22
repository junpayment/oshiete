package infrastructures

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/datastore"

	models "github.com/junpayment/oshiete/models/datastore"
)

type dataStoreClient interface {
	Put(ctx context.Context, key *datastore.Key, src interface{}) (*datastore.Key, error)
	Get(ctx context.Context, key *datastore.Key, dst interface{}) (err error)
}

type DataStoreClient struct {
	ctx             context.Context
	dataStoreClient dataStoreClient
}

func NewDataStoreClient(googleProjectName string) (*DataStoreClient, error) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, googleProjectName)
	if err != nil {
		return nil, fmt.Errorf(`dsClient, err := datastore.NewClient(c, googleProjectName): %w`, err)
	}
	return &DataStoreClient{
		ctx:             ctx,
		dataStoreClient: dsClient,
	}, nil
}

func (c *DataStoreClient) Save(key, answer string) error {
	data := &models.Answer{Body: answer}
	dataStoreKey := datastore.NameKey("Answer", key, nil)
	_, err := c.dataStoreClient.Put(c.ctx, dataStoreKey, data)
	if err != nil {
		return fmt.Errorf(`_, err := c.dataStoreClient.Put(ctx, dataStoreKey, data): %w`, err)
	}
	return nil
}

func (c *DataStoreClient) GetByKey(key string) (string, error) {
	data := &models.Answer{}
	dataStoreKey := datastore.NameKey("Answer", key, nil)
	err := c.dataStoreClient.Get(c.ctx, dataStoreKey, data)
	if err != nil && !errors.Is(err, datastore.ErrNoSuchEntity) {
		return "", fmt.Errorf(`err := c.dataStoreClient.Get(ctx, dataStoreKey, data): %w`, err)
	}
	return data.Body, nil
}
