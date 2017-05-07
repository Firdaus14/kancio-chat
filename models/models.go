package models

type Users struct {
	Username string `bson:"Username"`
	Nama     string `bson:"Nama"`
	Email    string `bson:"Email"`
	Password []byte `bson:"Password"`
	Jk       string `bson:"Jk"`
}
