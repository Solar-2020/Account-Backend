package models

type User struct {
	ID        int
	Email     string
	Name      string
	SureName  string
	AvatarURL string
}

type YandexUser struct {
	ID              string   `json:"id"`
	Login           string   `json:"login"`
	ClientID        string   `json:"client_id"`
	DisplayName     string   `json:"display_name"`
	RealName        string   `json:"real_name"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Sex             string   `json:"sex"`
	DefaultEmail    string   `json:"default_email"`
	Emails          []string `json:"emails"`
	DefaultAvatarID string   `json:"default_avatar_id"`
	IsAvatarEmpty   bool     `json:"is_avatar_empty"`
}
