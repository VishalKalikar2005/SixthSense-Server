package models
type User struct {
	ID       string `json:"_id" bsom:"_id"`
	Username string `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}
type UserLogin struct{
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}