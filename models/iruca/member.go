package iruca

import "time"

type Member struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Position  int       `json:"position"`
}
