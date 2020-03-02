package models

import (
	"fmt"
	"strings"
)

const (
	CommandEte       = "ete"
	CommandEru       = "eru"
	CommandIrukaList = "list"
)

// see https://github.com/gin-gonic/gin#bind-query-string-or-post-data
type SlackSlashCommandRequest struct {
	Token        string `form:"token"`
	Command      string `form:"command"`
	Text         string `form:"text"`
	ResponseUrl  string `form:"response_url"`
	TriggerId    string `form:"trigger_id"`
	UserId       string `form:"user_id"`
	UserName     string `form:"user_name"`
	TeamId       string `form:"team_id"`
	EnterpriseId string `form:"enterprise_id"`
	ChannelId    string `form:"channel_id"`
}

type SlackSlashCommandText struct {
	Command string
	Body1   string
	Body2   string
}

type SlackSlashCommandRequestOshi SlackSlashCommandRequest

func (s *SlackSlashCommandRequestOshi) ParseText() (*SlackSlashCommandText, error) {
	temp := strings.Split(s.Text, " ")
	if len(temp) < 2 {
		return nil, fmt.Errorf("not enough command num")
	}
	command := temp[0]
	if command != CommandEru && command != CommandEte {
		return nil, fmt.Errorf("invalid command: %s", command)
	}
	body1 := temp[1]
	var body2 string
	if len(temp) >= 3 {
		body2 = temp[2]
	}
	return &SlackSlashCommandText{
		Command: command,
		Body1:   body1,
		Body2:   body2,
	}, nil
}

type SlackSlashCommandRequestIruka SlackSlashCommandRequest

func (s *SlackSlashCommandRequestIruka) ParseText() (*SlackSlashCommandText, error) {
	temp := strings.Split(s.Text, " ")
	if len(temp) < 1 {
		return nil, fmt.Errorf("not enough command num")
	}
	command := temp[0]
	if command != CommandIrukaList {
		return nil, fmt.Errorf("invalid command: %s", command)
	}
	return &SlackSlashCommandText{
		Command: command,
	}, nil
}
