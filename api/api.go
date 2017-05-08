package api

import (
	"net/http"

	"github.com/DaanWillems/NanoRouter"
)

func StartApi(r *NanoRouter.Router) {
	r.NewRoute("GET", "/", home)
	r.NewRoute("GET", "/getPost/:id", getPost)
	r.NewRoute("GET", "/getPosts", getPosts)
}

func home(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	http.ServeFile(w, r, "../../view/home.html")
}
