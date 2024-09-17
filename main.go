package main

import (
	"fmt"
	"net/http"
	"os"
	"portfolio/internal/middleware"
	"portfolio/internal/template"
	"portfolio/internal/view"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	r := initializeRoutes()

	fmt.Printf("server is running on port %s\n", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		fmt.Println(err)
	}
}

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /static/", view.ServeStaticFiles)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})

	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/home" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, template.Home("home"))
	})

	return mux
}
