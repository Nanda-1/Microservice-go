package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	fmt.Println("Starting front end service on port 80")
	errr := http.ListenAndServe(":80", nil)
	if errr != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"D:/IBS Nanda/LearnMicroservices/front_end/cmd/web/templates/base.layout.gohtml",
		"D:/IBS Nanda/LearnMicroservices/front_end/cmd/web/templates/header.partial.gohtml",
		"D:/IBS Nanda/LearnMicroservices/front_end/cmd/web/templates/footer.partial.gohtml",
		// "D:/IBS Nanda/LearnMicroservices/front_end/cmd/web/templates/test.page.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("D:/IBS Nanda/LearnMicroservices/front_end/cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// Logging to check the paths being used
	for _, path := range templateSlice {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("Template file not found: %s\n", path)
		} else if err != nil {
			log.Printf("Error checking template file: %s, %v\n", path, err)
		}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
