package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wgarcia1309/go-twitter/middleware"
	"github.com/wgarcia1309/go-twitter/routers"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middleware.CheckDB(routers.NewUser)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil",
		middleware.CheckDB(
			middleware.ValidJWT(
				routers.GetProfile,
			),
		)).Methods("GET")
	router.HandleFunc("/modificarPerfil",
		middleware.CheckDB(
			middleware.ValidJWT(
				routers.UpdateProfile,
			),
		)).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
