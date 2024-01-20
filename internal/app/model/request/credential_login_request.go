package request

type CredentialLoginRequest struct {
	Username string `form:"username" json:"username" bson:"username"`
	Password string `form:"password" json:"password" bson:"password"`
}
