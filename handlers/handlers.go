package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/valianx/clon-twitter/middlew"
	"log"
	"net/http"
	"os"
)
//configuracion del servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", ChequeoDB(routers.Registro)).Methods("POST")



	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
