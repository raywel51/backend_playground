package model

type CredentialRegisterRequest struct {
	Username string `form:"username" json:"username" bson:"username"`
	Password string `form:"password" json:"password" bson:"password"`
	Email    string `form:"email" json:"email" bson:"email"`
}
