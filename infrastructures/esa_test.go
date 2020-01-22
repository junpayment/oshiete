package infrastructures

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hiroakis/esa-go/response"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/infrastructures/mocks"
)

func TestEsaClient_GetListByKey(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		esaClient := mocks.NewMockesaClient(ctrl)
		esaClient.EXPECT().SetPage(gomock.Any()).Return()
		esaClient.EXPECT().SetQuery(gomock.Any()).Return()
		esaClient.EXPECT().GetPosts().Return(response.Posts{}, nil)
		client := &EsaClient{
			esaClient: esaClient,
		}
		res, err := client.GetListByKey("test key")
		assert.Nil(t, err)
		assert.Len(t, res.List, 0)
	})
}
