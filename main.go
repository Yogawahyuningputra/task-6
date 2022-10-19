package main

// import "fmt"

// func main()	{
// 	var name string = "Yoga"

// 	var age int = 25
// 	fmt.Println("hello world")
// 	fmt.Println(name)
// 	fmt.Println(age) ()
// }

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")) ))
		
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/index", index).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/detailproject", detailProject).Methods("GET")
	route.HandleFunc("/project", addProject).Methods("POST")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return

	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func detailProject (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/detailproject.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	 tmpl.Execute(w, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte ("message : " + err.Error()))
		return

	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
		
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)

}
func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name project : " + r.PostForm.Get("input-project"))
	fmt.Println("star date : " + r.PostForm.Get("input-star-date"))
	fmt.Println("end date : " + r.PostForm.Get("input-end-date"))
	fmt.Println("content : " + r.PostForm.Get("input-content"))
	fmt.Println("technology : " + r.PostForm.Get("checkbox1"))
	fmt.Println("technology : " + r.PostForm.Get("checkbox2"))
	fmt.Println("technology : " + r.PostForm.Get("checkbox3"))
	fmt.Println("technology : " + r.PostForm.Get("checkbox4"))
	fmt.Println("upload file : " + r.PostForm.Get("input-image"))

	http.Redirect(w, r, "/project", http.StatusMovedPermanently)

}