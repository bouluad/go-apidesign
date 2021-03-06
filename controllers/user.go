package controllers

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/go-apidesign/services"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (this UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService()
	json, _ := json.Marshal(userService.GetUsers())
	fmt.Fprint(w, string(json))
}
