package main


import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Name": "Personal Web",
}
type Project struct {
	Name		string
	Post_date	string
	Author		string
	Content		string
	Technology	string

	
}

var Projects =[]Project{

	{
	Name		:"Dumbways Web Apps",
	Post_date	:"20/10/2022 | 20/12/2022",
	Author		:"Yoga",
	Content		:"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
	Technology	:"Golang",
	},
	


}


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
	route.HandleFunc("deleteProject/{index}", delProject).Methods("GET")

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

func detailProject (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/detailproject.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var DetailProject = Project{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range Projects {
		if index == i {
			DetailProject = Project {
				Name: data.Name,
				Post_date: data.Post_date,
				Content: data.Content,
				Technology: data.Technology,
				Author: data.Author,

			}
		}
		
	}
	data := map[string]interface{}{
		"Project": DetailProject,
	}
	

	w.WriteHeader(http.StatusOK)
	 tmpl.Execute(w, data)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	respData := map[string]interface{}{
		"Projects": Projects,
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)

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


	var name = r.PostForm.Get("input-project")
	var content = r.PostForm.Get("input-content")
	
	var newProject = Project{

		Name :	name,
		Content: content,
		Author: "Yoga",
		Post_date: time.Now().String(),
		Technology: "Golang",

	}
	Projects = append(Projects, newProject)


	http.Redirect(w, r, "/project", http.StatusMovedPermanently)

}

func delProject (w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	Projects = append(Projects[:index], Projects[index+1:]...)
	fmt.Println(Projects)

	http.Redirect(w, r, "/project", http.StatusFound)
}