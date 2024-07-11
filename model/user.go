package model

type User struct {
	_id        string
	username   string
	email      string
	password   string
	birthdate  string
	createDate string
	updateDate string
	isActive   bool
	isDeleted  bool
}

type UserProfile struct {
	user      User
	followers int
	following int
	posts     []string
}
