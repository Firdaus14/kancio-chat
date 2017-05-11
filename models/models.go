package models

import "time"

type Users struct {
	Username string `bson:"Username"`
	Nama     string `bson:"Nama"`
	Email    string `bson:"Email"`
	Password []byte `bson:"Password"`
	Jk       string `bson:"Jk"`
}

type Home struct {
	Nama  string
	Index []Users
}

type Session struct {
	Username     string
	LastActivity time.Time
}
