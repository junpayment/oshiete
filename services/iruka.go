package services

import (
	"fmt"

	"github.com/junpayment/oshiete/models/iruca"
	"github.com/junpayment/oshiete/models/iruka"
)

type IrucaClient interface {
	GetMembers() ([]*iruca.Member, error)
}

type IrukaService struct {
	IrucaClient IrucaClient
}

func (s *IrukaService) GetList() ([]*iruka.State, error) {
	members, err := s.IrucaClient.GetMembers()
	if err != nil {
		return nil, fmt.Errorf(`members, err := s.IrucaClient.GetMembers(): %w`, err)
	}
	var states []*iruka.State
	for _, member := range members {
		states = append(states, &iruka.State{
			Name:    member.Name,
			Status:  member.Status,
			Message: member.Message,
		})
	}
	return states, nil
}
