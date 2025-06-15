package entity

type User struct {
	Id          string `bson:"_id,omitempty"`
	UserName    string `bson:"user_name"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	PhoneNumber string `bson:"phone_number"`
}

func NewUser(id, userName, email, password, phonenumber string) *User {
	return &User{
		Id:          id,
		UserName:    userName,
		Email:       email,
		Password:    password,
		PhoneNumber: phonenumber,
	}
}
