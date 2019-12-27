package service

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo!")
	fmt.Println("Endpoint Hit: homePage")
}
