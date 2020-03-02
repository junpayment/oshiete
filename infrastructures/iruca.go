package infrastructures

import (
	"encoding/json"
	"fmt"
	"github.com/junpayment/oshiete/models/iruca"
	"io/ioutil"
	"net/http"
)

const (
	IrucaBaseURL = `https://iruca.co/api/rooms`
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type IrucaClient struct {
	httpClient httpClient
	roomId     string
	token      string
}

func NewIrucaClient(roomId, token string) *IrucaClient {
	return &IrucaClient{
		httpClient: http.DefaultClient,
		roomId:     roomId,
		token:      token,
	}
}

func (c *IrucaClient) GetMembers() ([]*iruca.Member, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(`%s/%s/%s`, IrucaBaseURL, c.roomId, "members"), nil)
	if err != nil {
		return nil, fmt.Errorf(
			`req, err := http.NewRequest(http.MethodGet, path.Join(IrucaBaseURL, c.roomId, "members"), nil): %w`,
			err)
	}
	req.Header.Set(`X-Iruca-Token`, c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(`res, err := c.httpClient.Do(req): %w`, err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(`b, err := ioutil.ReadAll(res.Body): %w`, err)
	}
	var members []*iruca.Member
	err = json.Unmarshal(b, &members)
	if err != nil {
		return nil, fmt.Errorf(`err := json.Unmarshal(b, members): %w`, err)
	}

	return members, nil
}
