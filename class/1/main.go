package main

import (
	"html/template"
	"log"
	"net/http"
)

//var (
//	tmpl *template.Template
//)

// Fungsi template untuk memformat tanggal
//var funcMap = template.FuncMap{
//	"formatDate": func(t time.Time) string {
//		return t.Format("02 January 2006")
//	},
//}

//func handler(w http.ResponseWriter, r *http.Request) {
//	data := struct {
//		Title string
//	}{
//		Title: "Halaman Utama",
//	}
//
//	// Menggunakan template yang telah di-cache
//	err := tmpl.ExecuteTemplate(w, "index.html", data)
//	if err != nil {
//		log.Fatalf("Error executing template: %v", err)
//	}
//}

//func handler(w http.ResponseWriter, r *http.Request) {
//	// Data yang diperlukan untuk template
//	data := struct {
//		Title string
//		Date  time.Time
//	}{
//		Title: "Halaman Utama",
//		Date:  time.Now(),
//	}
//
//	// Parse template dengan funcMap
//	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
//	if err != nil {
//		log.Fatalf("Error parsing template: %v", err)
//	}
//
//	// Execute template
//	err = tmpl.Execute(w, data)
//	if err != nil {
//		log.Fatalf("Error executing template: %v", err)
//	}
//}

func handler(w http.ResponseWriter, r *http.Request) {
	// Data yang diperlukan untuk layout dan halaman spesifik
	data := struct {
		Title string
	}{
		Title: "Halaman Utama",
	}

	// Parse layout template dan template-template lainnya
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute layout template
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	// Data yang diperlukan untuk layout dan halaman spesifik
	data := struct {
		Title string
	}{
		Title: "Halaman Utama",
	}

	// Parse layout template dan template-template lainnya
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute layout template
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/tentang", handler2)
	http.HandleFunc("/", handler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
