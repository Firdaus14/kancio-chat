package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kanciogo/kancio-chat/models"
	"golang.org/x/crypto/bcrypt"
)

func connect() *mgo.Session {
	var session, err = mgo.Dial("172.17.0.2")
	if err != nil {
		os.Exit(0)
	}
	return session
}

type Controllers struct {
	tpl *template.Template
}

func NewControlers(t *template.Template) *Controllers {
	return &Controllers{t}
}

func (c Controllers) Daftar(w http.ResponseWriter, r *http.Request) {
	sessions := connect()
	defer sessions.Close()
	collection := sessions.DB("kancio").C("user")
	var data models.Users

	if r.Method == http.MethodPost {
		Username := r.FormValue("Username")
		Nama := r.FormValue("Nama")
		Email := r.FormValue("Email")
		Password := r.FormValue("Password")
		Jk := r.FormValue("Jk")
		bs, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Gak bisa generate password", http.StatusInternalServerError)
			return
		}
		data = models.Users{Username, Nama, Email, bs, Jk}
		fmt.Println(data)
		simpan := collection.Insert(data)
		if simpan != nil {
			http.Error(w, "Gak bisa nyimpan data", http.StatusInternalServerError)
			return
		}
	}
	c.tpl.ExecuteTemplate(w, "daftar.html", data)
}

func (c Controllers) Login(w http.ResponseWriter, r *http.Request) {
	sessions := connect()
	defer sessions.Close()
	collection := sessions.DB("kancio").C("user")
	var data models.Users
	if r.Method == http.MethodPost {
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")
		var selector = bson.M{"Username": Username}
		collection.Find(selector).One(&data)
		err := bcrypt.CompareHashAndPassword(data.Password, []byte(Password))
		if err != nil {
			http.Error(w, "Username atau password salah", http.StatusForbidden)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	c.tpl.ExecuteTemplate(w, "login.html", data)
}
