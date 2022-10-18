package main

import (
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")            // Create a template.
	t, _ = t.ParseFiles("tmpl/welcome.html", nil) // Parse template file.
	user := map[string]string{"a": "a"}           // Sample info to embed
	t.Execute(w, user)                            // merge.
}

// {{.}} refers to the current object
// {{.a}} refers to the a field in current object

// {{with .}}…{{end}}

// {{range .}}…{{end}}

// {{.| html}} pipe the object in . to the html function

// {{with $x := "output" | printf "%q"}}{{$x}}{{end}}

// All template functions must be placed in a funcmap()

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}

// func main() {

// }
