package main

import (
	"fmt"
	authcontroller "go_auth/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server jalan di: http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
