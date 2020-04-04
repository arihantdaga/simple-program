package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/arihantdaga/simple-program/controllers"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Print(something)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Finduser("test", "sdad")
	Anything()
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	Connect()
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/learning", controllers.Learning)

	log.Fatal(http.ListenAndServe(":8080", router))
}
