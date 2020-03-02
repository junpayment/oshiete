package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/models/iruca"
	"github.com/junpayment/oshiete/services/mocks"
)

func TestIrukaService_GetList(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		irucaClient := mocks.NewMockIrucaClient(ctrl)
		irucaClient.EXPECT().GetMembers().Return([]*iruca.Member{
			{
				Name:    "test",
				Status:  "test",
				Message: "test",
			},
		}, nil)
		service := &IrukaService{IrucaClient: irucaClient}
		res, err := service.GetList()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
