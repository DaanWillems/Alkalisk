package api

import (
	"Alkalisk/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getComments(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	ur := model.PostRepository{}
	id, _ := strconv.ParseInt(vars["id"], 10, 0)
	p := ur.GetCommentsFromPage(int(id))

	response, err := json.Marshal(p)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, string(response))
}

func newComment(w http.ResponseWriter, r *http.Request, vars map[string]string) {

	CommentRepository := model.CommentRepository{}

	decoder := json.NewDecoder(r.Body)

	req := struct {
		Id      string `json:"id"`
		Content string `json:"content"`
	}{}

	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	id, _ := strconv.Atoi(req.Id)
	comment := CommentRepository.New(req.Content, id)
	response, err := json.Marshal(comment)
	fmt.Fprintf(w, string(response))
}
