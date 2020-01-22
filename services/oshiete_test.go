package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/models"
	"github.com/junpayment/oshiete/services/mocks"
)

func TestOshiete_Eru(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		dsClient := mocks.NewMockDataStoreClient(ctrl)
		dsClient.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		service := &Oshiete{
			DataStoreClient: dsClient,
		}

		err := service.Eru("test key", "test answer")
		assert.Nil(t, err)
	})
}

func TestOshiete_Ete(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		dsClient := mocks.NewMockDataStoreClient(ctrl)
		esaClient := mocks.NewMockEsaClient(ctrl)

		dsClient.EXPECT().GetByKey(gomock.Any()).Return("this is test, answer", nil)
		esaClient.EXPECT().GetListByKey(gomock.Any()).Return(&models.EsaAnswer{
			List: []*struct {
				Name        string
				Description string
				URL         string
			}{
				{
					Name:        "this is test, esa answer",
					Description: "this is test, esa answer description",
					URL:         "this is test, esa answer url",
				},
			},
		}, nil)

		service := &Oshiete{
			DataStoreClient: dsClient,
			EsaClient:       esaClient,
		}

		answer, err := service.Ete("test key")
		assert.Nil(t, err)
		assert.Equal(t, &models.EteAnswer{
			Answer: "this is test, answer",
			EsaAnswer: &models.EsaAnswer{
				List: []*struct {
					Name        string
					Description string
					URL         string
				}{
					{
						Name:        "this is test, esa answer",
						Description: "this is test, esa answer description",
						URL:         "this is test, esa answer url",
					},
				},
			},
		}, answer)
	})
}
