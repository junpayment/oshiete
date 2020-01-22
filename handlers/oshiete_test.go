package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/handlers/mocks"
	"github.com/junpayment/oshiete/middlewares"
	"github.com/junpayment/oshiete/models"
)

func TestOshiete_Do(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("command eru", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		oshieteService := mocks.NewMockOshieteService(ctrl)
		oshieteService.EXPECT().Eru(gomock.Any(), gomock.Any()).Return(nil)
		handler := &Oshiete{
			OshieteService: oshieteService,
		}

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		form := url.Values{}
		form.Add("text", "eru test_key_string test_answer_string")
		body := strings.NewReader(form.Encode())
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		handler.Do(c)

		err, exists := c.Get(middlewares.KeyError)
		assert.Nil(t, err)
		assert.False(t, exists)
		resText, exists := c.Get(middlewares.KeyResponseText)
		assert.True(t, exists)
		assert.Equal(t, "教えてくれてありがとうございます！", resText)
	})

	t.Run("command ete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		oshieteService := mocks.NewMockOshieteService(ctrl)
		templateService := mocks.NewMockTemplateService(ctrl)
		oshieteService.EXPECT().Ete(gomock.Any()).Return(&models.EteAnswer{}, nil)
		templateService.EXPECT().OutEte(gomock.Any()).Return("this is test", nil)
		handler := &Oshiete{
			OshieteService:  oshieteService,
			TempleteService: templateService,
		}

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		form := url.Values{}
		form.Add("text", "ete test_key_string")
		body := strings.NewReader(form.Encode())
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		handler.Do(c)

		err, exists := c.Get(middlewares.KeyError)
		assert.Nil(t, err)
		assert.False(t, exists)
		resText, exists := c.Get(middlewares.KeyResponseText)
		assert.True(t, exists)
		assert.Equal(t, "this is test", resText)
	})
}
