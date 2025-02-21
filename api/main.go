package handler

import (
  "fmt"
  "net/http"
  "html/template"
)

func main() {
  http.HandleFunc("/", HomeHandler)
  http.HandleFunc("/repeat", RepeatHandler)

  fmt.Println("Server started at port 8080")
  http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    input := r.FormValue("text")
    data := PageData{Text: input}
    tmpl := template.Must(template.ParseFiles("static/result.html"))
    tmpl.Execute(w, data)
  }
}

