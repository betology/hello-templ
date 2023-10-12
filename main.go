package main

import (
	"log"
	"net/http"
)

func main() {
        http.Handle("/hello", HelloHandler{})
        http.ListenAndServe(":8000", nil)
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // Get the data.
  name := r.URL.Query().Get("name")

  // Create the components.
  body := Greeter(name)
  page := Page("hello", body)

  // Render.
  err := page.Render(r.Context(), w)
  if err != nil {
          log.Println("error", err)
  }
}
