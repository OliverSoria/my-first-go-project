package service

import (
	"../model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {

	articles := []model.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(articles)
}
