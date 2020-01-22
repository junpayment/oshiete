package services

import (
	"fmt"
	"github.com/junpayment/oshiete/models"
)

type EsaClient interface {
	GetListByKey(key string) (*models.EsaAnswer, error)
}

type DataStoreClient interface {
	Save(key, answer string) error
	GetByKey(key string) (string, error)
}

type Oshiete struct {
	EsaClient       EsaClient
	DataStoreClient DataStoreClient
}

func (s *Oshiete) Eru(key, answer string) error {
	return s.DataStoreClient.Save(key, answer)
}

func (s *Oshiete) Ete(key string) (*models.EteAnswer, error) {
	answer, err := s.DataStoreClient.GetByKey(key)
	if err != nil {
		return nil, fmt.Errorf(`answer, err := s.DataStoreClient.GetByKey(key): %w`, err)
	}
	esaAnswer, err := s.EsaClient.GetListByKey(key)
	if err != nil {
		return nil, fmt.Errorf(`esaAnswer, err := s.EsaClient.GetListByKey(key): %w`, err)
	}
	return &models.EteAnswer{
		Answer:    answer,
		EsaAnswer: esaAnswer,
	}, nil
}
