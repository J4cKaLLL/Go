package main

import (
	"fmt"
	"net/http"

	"examples.com/packages/util"

	"github.com/gorilla/mux"
)

func main() {
	greeting := fmt.Sprintf("Hello, %s", "James")
	fmt.Println(greeting)

	fmt.Printf("Length of greeting is %d\n", util.StringLength(greeting))
	fmt.Printf(util.GetGreeting())
	r := mux.NewRouter()
	http.ListenAndServe(":9000", r)
}
