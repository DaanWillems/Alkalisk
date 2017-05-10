package api

import (
	"Alkalisk/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getPost(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	ur := model.PostRepository{}
	id, _ := strconv.ParseInt(vars["id"], 10, 0)
	p := ur.Get(int(id))

	response, err := json.Marshal(p)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, string(response))
}

func getPosts(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	postRepository := model.PostRepository{}
	posts := postRepository.All(0, 0)
	response, err := json.Marshal(posts)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, string(response))
}

func newPost(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	PostRepository := model.PostRepository{}

	decoder := json.NewDecoder(r.Body)

	req := struct {
		Id      string `json:"id"`
		Title   string `json:string`
		Content string `json:"content"`
	}{}

	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	comment := PostRepository.New(req.Title, req.Content)
	response, err := json.Marshal(comment)
	fmt.Fprintf(w, string(response))
}
