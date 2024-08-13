package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"rest_api_sklad_project/pkg/handlers"

	"github.com/gorilla/mux"

	// swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"

)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sklad REST API!")
	fmt.Println("SKLAD REST API")
}

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/", homePage)
	r.HandleFunc("/goods", handlers.GetAllGoods).Methods(http.MethodGet)
	r.HandleFunc("/goods/{id}", handlers.GetGoods).Methods(http.MethodGet)
	r.HandleFunc("/goods", handlers.AddGood).Methods(http.MethodPost)
	r.HandleFunc("/goods/{id}", handlers.UpdateGoods).Methods(http.MethodPut)
	r.HandleFunc("/goods/{id}", handlers.DeleteGoods).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}