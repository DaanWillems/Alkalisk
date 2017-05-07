package api

import (
	"Alkalisk/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func getTopics(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	TopicRepository := model.TopicRepository{}

	response, err := json.Marshal(TopicRepository.All(0, 0))

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, string(response))
}
