package model

type Webinar struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkMeet    string `json:"link_meet"`
	HostID      int    `json:"host_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
