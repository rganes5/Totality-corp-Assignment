package models

type UserData struct {
	UserId    int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

type NotFoundList struct {
	UsersNotFound []int32
}
