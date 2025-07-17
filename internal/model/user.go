package model

type User struct {
	ID             string `json:"id"`
	DeviceID       string `json:"device_id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	LastLogin      string `json:"last_login"`
	Active         bool   `json:"active"`
}
