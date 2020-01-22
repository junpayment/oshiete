package infrastructures

import (
	"context"
	models "github.com/junpayment/oshiete/models/datastore"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/infrastructures/mocks"
)

func TestDataStoreClient_Save(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		dsClient := mocks.NewMockdataStoreClient(ctrl)
		dsClient.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(&datastore.Key{}, nil)
		client := &DataStoreClient{
			dataStoreClient: dsClient,
		}
		err := client.Save("test key", "test answer")
		assert.Nil(t, err)
	})
}

func TestDataStoreClient_GetByKey(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		dsClient := mocks.NewMockdataStoreClient(ctrl)
		dsClient.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, key *datastore.Key, src interface{}) error {
				src.(*models.Answer).Body = "this is test body"
				return nil
			})
		client := &DataStoreClient{
			dataStoreClient: dsClient,
		}
		res, err := client.GetByKey("test key")
		assert.Nil(t, err)
		assert.Equal(t, "this is test body", res)
	})
}
