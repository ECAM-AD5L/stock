package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/stock", CreateStock).Methods("POST")
	router.HandleFunc("/stock", ModifyStock).Methods("PUT")
	router.HandleFunc("/stock/{id}", GetStock).Methods("GET")

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println("Error parsing the DEBUG env variable")
	}
	if debug {
		err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pathTemplate, err := route.GetPathTemplate()
			if err == nil {
				fmt.Print("ROUTE:", pathTemplate)
			}
			methods, err := route.GetMethods()
			if err == nil {
				fmt.Println(" Methods:", strings.Join(methods, ","))
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
