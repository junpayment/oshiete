package handlers

import (
	"github.com/junpayment/oshiete/models"
	"github.com/junpayment/oshiete/models/iruka"
)

type TemplateService interface {
	OutEte(answer *models.EteAnswer) (string, error)
	OutIruka(states []*iruka.State) (string, error)
}
