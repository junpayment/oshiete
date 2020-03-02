package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"

	"github.com/junpayment/oshiete/middlewares"
	"github.com/junpayment/oshiete/models"
	"github.com/junpayment/oshiete/models/iruka"
)

type IrukaService interface {
	GetList() ([]*iruka.State, error)
}

type IrukaHandler struct {
	IrukaService    IrukaService
	TemplateService TemplateService
}

type statesSorted []*iruka.State

func (s statesSorted) Len() int {
	return len(s)
}
func (s statesSorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s statesSorted) Less(i, j int) bool {
	return s[i].Status < s[j].Status
}

func (h *IrukaHandler) Do(c *gin.Context) {
	body := &models.SlackSlashCommandRequestIruka{}
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

	switch textBody.Command {
	case models.CommandIrukaList:
		states, err := h.IrukaService.GetList()
		if err != nil {
			c.Set(middlewares.KeyError, err)
			return
		}
		var statesS statesSorted = states
		sort.Sort(statesS)

		res, err := h.TemplateService.OutIruka(statesS)
		if err != nil {
			c.Set(middlewares.KeyError, err)
			return
		}
		c.Set(middlewares.KeyResponseText, res)
		return
	default:
		c.Set(middlewares.KeyError, fmt.Errorf("不正なコマンドです"))
		return
	}
}
