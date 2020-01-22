package infrastructures

import (
	"fmt"

	"github.com/hiroakis/esa-go"
	"github.com/hiroakis/esa-go/response"

	"github.com/junpayment/oshiete/models"
)

type esaClient interface {
	SetQuery(query string)
	SetPage(page int)
	GetPosts() (response.Posts, error)
}

type EsaClient struct {
	// 使ってるesa clientがqueryを初期化する手段がないけどSetQuery()は内部のQueryを初期化する実装のためよろしくないけどこのままで行く
	esaClient esaClient
}

func NewEsaClient(apiKey, team string) *EsaClient {
	return &EsaClient{
		esaClient: esa.NewEsaClient(apiKey, team),
	}
}

func (c *EsaClient) GetListByKey(key string) (*models.EsaAnswer, error) {
	cli := c.esaClient
	cli.SetPage(1)
	cli.SetQuery("keyword: " + key)
	posts, err := cli.GetPosts()
	if err != nil {
		return nil, fmt.Errorf(`posts, err := cli.GetPosts(): %w`, err)
	}
	var list []*struct {
		Name        string
		Description string
		URL         string
	}
	for _, post := range posts.Posts {
		str := &struct {
			Name        string
			Description string
			URL         string
		}{
			Name:        post.Name,
			Description: post.BodyMd,
			URL:         post.Url,
		}
		list = append(list, str)
	}
	return &models.EsaAnswer{
		List: list,
	}, nil
}
