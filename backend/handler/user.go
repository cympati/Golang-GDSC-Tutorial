package handler

import (
	"cympati/Golang-GDSC-Tutorial/service"
	"net/http"
)

type userHandler struct {
	service service.UserService // interface หรือ "Port" จาก Service
}

func NewUserHandler(userService service.UserService) userHandler { // รับ "Adapter" ของ Service
	return userHandler{service: userService}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignUp("", "")
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignIn("", "")
}

func (h userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.ListUsers()
}
