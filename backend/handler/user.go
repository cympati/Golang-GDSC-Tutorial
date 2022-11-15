package handler

import (
	"cympati/Golang-GDSC-Tutorial/service"
	"cympati/Golang-GDSC-Tutorial/types"
	"cympati/Golang-GDSC-Tutorial/utils"
	"encoding/json"
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
	//h.service.SignUp("", "")

	// Check if the request method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	var response []byte
	var body types.SignUp
	err := utils.Parse(r, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call signup service
	token, base64, err := h.service.SignUp(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response
	response, _ = json.Marshal(map[string]any{"success": true, "token": token, "image": base64})
	w.Write(response)
	//fmt.Fprint(w, response)
	return
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignIn("", "")
}

func (h userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.ListUsers()
}
