package entities

type FirebaseUser struct {
	Firstname string `json:"firstname" validate:"required" bson:"firstname"`
	Lastname string `json:"lastname" validate:"required" bson:"lastname"`
	Email string `json:"email" validate:"required" bson:"email"`
	Password string `json:"password" bson:"password"`
	AuthId string `json:"auth_id" bson:"auth_id"`
}

type FirebaseToken struct {
	Email string `json:"email"`
	IDToken string `json:"id_token"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
	IDToken string `json:"id_token"`
}
