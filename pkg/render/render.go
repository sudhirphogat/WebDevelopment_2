package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemp(w http.ResponseWriter, tmpl string) {
	parseTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("this is an error", err)
	}

	ts := parseTemplate.Execute(w, nil)

	if ts != nil {
		fmt.Println("this ts error", ts)
		return
	}
}
