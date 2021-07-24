package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/xavimg/TweetGO/middlewares"
	"github.com/xavimg/TweetGO/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar el Server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
