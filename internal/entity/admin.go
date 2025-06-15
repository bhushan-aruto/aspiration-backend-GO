package entity

type Admin struct {
	Id          string `bson:"_id"`
	AdminName   string `bson:"admin_name"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	Password    string `bson:"password"`
}

func NewAdmin(id, adminName, email, phoneNumber, password string) *Admin {
	return &Admin{
		Id:          id,
		AdminName:   adminName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}

}
