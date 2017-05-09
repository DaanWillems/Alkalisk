package api

import (
	"fmt"
	"net/http"

	"github.com/DaanWillems/NanoRouter"
)

func StartApi(r *NanoRouter.Router) {
	r.NewRoute("GET", "/", home)
	r.NewRoute("GET", "/getComments/:id", getComments)
	r.NewRoute("POST", "/postComment/:id", postComment)
	r.NewRoute("GET", "/getPost/:id", getPost)
	r.NewRoute("GET", "/getPosts", getPosts)
	r.SetNotFoundRoute(home)
}

func test(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	fmt.Println("test123")
	fmt.Fprintf(w, vars["id"])
}

func home(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	fmt.Println("test")
	http.ServeFile(w, r, "../../view/home.html")
}
