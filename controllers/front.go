package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserConstroller()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
