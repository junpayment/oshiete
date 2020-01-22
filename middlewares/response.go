package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/junpayment/oshiete/models"
)

const (
	KeyResponseText = "res_text"
	KeyError        = "err"
)

func Response(c *gin.Context) {
	c.Next()
	err, exists := c.Get(KeyError)
	if exists == true {
		res := &models.SlackSlashCommandResponse{
			ResponseType: models.ResponseTypeEphemeral,
			Text:         "エラーです: " + err.(error).Error(),
		}
		c.JSON(http.StatusOK, res)
		return
	}
	resText, _ := c.Get(KeyResponseText)
	res := &models.SlackSlashCommandResponse{
		ResponseType: models.ResponseTypeInChannel,
		Text:         resText.(string),
	}
	c.JSON(http.StatusOK, res)
}
