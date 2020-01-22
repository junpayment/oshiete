package models

const (
	ResponseTypeEphemeral = "ephemeral"
	ResponseTypeInChannel = "in_channel"
)

type SlackSlashCommandResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
