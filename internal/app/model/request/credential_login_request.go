package request

import "github.com/go-playground/validator/v10"

type CredentialLoginRequest struct {
	Username string `form:"username" json:"username" bson:"username"`
	Password string `form:"password" json:"password" bson:"password"`
}

func (c *CredentialLoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
