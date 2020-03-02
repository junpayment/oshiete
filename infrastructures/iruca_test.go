package infrastructures

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/junpayment/oshiete/infrastructures/mocks"
	"github.com/junpayment/oshiete/models/iruca"
)

func TestIrucaClient_GetMembers(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		httpClient := mocks.NewMockhttpClient(ctrl)
		var members []*iruca.Member
		members = []*iruca.Member{
			{
				ID:        1,
				RoomID:    1,
				Name:      "test",
				Status:    "test",
				Message:   "test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Position:  1,
			},
		}
		b, _ := json.Marshal(members)
		body := ioutil.NopCloser(bytes.NewBuffer(b))
		httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
			Body: body,
		}, nil)
		client := &IrucaClient{httpClient: httpClient}
		res, err := client.GetMembers()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
