package model

type Error struct {
	Code		int
	Message		string
	Description	string
}

type User struct {
	ID		string
	Pass	string
}
