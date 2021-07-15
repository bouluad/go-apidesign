package api

import (
	"net/http"

	"log"

	"github.com/go-apidesign/routers"
)

func Run() {
	userRouter := routers.NewUserRouter()
	log.Fatal(http.ListenAndServe(":8080", userRouter.GetRouter()))
}
