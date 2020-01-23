package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/junpayment/oshiete/middlewares"
	"github.com/junpayment/oshiete/models"
)

type OshieteService interface {
	Eru(key, answer string) error
	Ete(key string) (*models.EteAnswer, error)
}

type TemplateService interface {
	OutEte(answer *models.EteAnswer) (string, error)
}

type Oshiete struct {
	OshieteService  OshieteService
	TempleteService TemplateService
}

func (h *Oshiete) Do(c *gin.Context) {
	body := &models.SlackSlashCommandRequest{}
	err := c.ShouldBind(body)
	if err != nil {
		c.Set(middlewares.KeyError, err)
		return
	}
	textBody, err := body.ParseText()
	if err != nil {
		c.Set(middlewares.KeyError, err)
		return
	}

	if textBody.Command == models.CommandEru {
		err = h.OshieteService.Eru(textBody.Body1, textBody.Body2)
		if err != nil {
			c.Set(middlewares.KeyError, err)
			return
		}
		c.Set(middlewares.KeyResponseText, textBody.Body1+" は "+textBody.Body2+"\nなんですね！教えてくれてありがとうございます！")
		return
	} else if textBody.Command == models.CommandEte {
		answer, err := h.OshieteService.Ete(textBody.Body1)
		if err != nil {
			c.Set(middlewares.KeyError, err)
			return
		}
		res, err := h.TempleteService.OutEte(answer)
		if err != nil {
			c.Set(middlewares.KeyError, err)
			return
		}
		c.Set(middlewares.KeyResponseText, res)
		return
	} else {
		c.Set(middlewares.KeyError, fmt.Errorf("不正なコマンドです"))
		return
	}
}
