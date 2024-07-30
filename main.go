package main

import (
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	Title       string
	Description string
	Image       string
	Link        string
}

type Skill struct {
	Name string
	Icon string
}

type PageData struct {
	Title    string
	Projects []Project
	Skills   []Skill
}

func main() {
	projects := []Project{
		{"Let's Move Fitness Studio", "Comprehensive Web Application Development for local business with ASP.net core MVC and mysql", "/static/project1.jpg", "#"},
		{"Native American Arts and Crafts Market", "Empowering Indigenous Artistry with a Dynamic Website for NAACM", "/static/project2.jpg", "#"},
		{"Project 3", "Description for project 3", "/static/project3.jpg", "#"},
	}

	skills := []Skill{
		{"HTML5", "/static/icon-html5.png"},
		{"CSS3", "/static/icon-css3.png"},
		{"JavaScript", "/static/icon-js.png"},
		{"React", "/static/icon-react.png"},
		{"Node.js", "/static/icon-node.png"},
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/index.html",
	))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:    "Developer Portfolio",
			Projects: projects,
			Skills:   skills,
		}
		err := tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		}
	})

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
