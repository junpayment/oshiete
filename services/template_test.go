package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/models"
)

func TestTemplete_OutEte(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		service := &Templete{}
		a := &models.EteAnswer{
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
		}
		res, err := service.OutEte(a)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})
}
