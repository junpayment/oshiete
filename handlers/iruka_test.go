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
	"github.com/junpayment/oshiete/models/iruka"
)

func TestIrukaHandler_Do(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("command list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		irukaService := mocks.NewMockIrukaService(ctrl)
		irukaService.EXPECT().GetList().Return([]*iruka.State{
			{
				Name:    "test",
				Status:  "test",
				Message: "test",
			},
			{
				Name:    "test3",
				Status:  "test",
				Message: "test",
			},
			{
				Name:    "test1",
				Status:  "test",
				Message: "test",
			},
		}, nil)
		templateService := mocks.NewMockTemplateService(ctrl)
		templateService.EXPECT().OutIruka(gomock.Any()).Return("this is test", nil)
		handler := &IrukaHandler{
			IrukaService:    irukaService,
			TemplateService: templateService,
		}
		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		form := url.Values{}
		form.Add("text", "list")
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
