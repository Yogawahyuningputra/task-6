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
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello world !"))

	})

	
 	fmt.Println("Server running on port 5000")
 	http.ListenAndServe("localhost:5000", route)
}