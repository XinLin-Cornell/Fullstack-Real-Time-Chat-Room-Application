package responses

import "chat/models"

type SignUpResponse struct {
	User     models.User `json:"User"`
	JwtToken string      `json:"Token"`
}
