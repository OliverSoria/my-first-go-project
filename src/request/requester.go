package request

import (
	"../service"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/hello", service.HomePage)
	http.HandleFunc("/articles", service.GetArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
