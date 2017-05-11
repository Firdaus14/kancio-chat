package main

import (
	"html/template"
	"net/http"

	"github.com/kanciogo/kancio-chat/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	c := controllers.NewControlers(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/daftar", c.Daftar)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/home", c.Home)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
