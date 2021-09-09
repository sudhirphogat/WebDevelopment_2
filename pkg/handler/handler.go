package handler

import (
	"fmt"
	"net/http"

	"github.com/sudhirphogat/WebDevelopment_2/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is home Page")
	render.RenderTemp(w, "home.page.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is about Page")
	render.RenderTemp(w, "about.page.html")
}
