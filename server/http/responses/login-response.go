package responses

import "chat/models"

type LoginResponse struct {
	User     models.User `json:"User"`
	JwtToken string      `json:"Token"`
}
