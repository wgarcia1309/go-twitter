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
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")

	userRouter := router.PathPrefix("user").Subrouter()
	userRouter.HandleFunc("/", middleware.CheckDB(routers.NewUser)).Methods("POST")
	userRouter.HandleFunc("/", middleware.CheckDB(middleware.ValidJWT(routers.GetProfile))).Methods("GET")
	userRouter.HandleFunc("/", middleware.CheckDB(middleware.ValidJWT(routers.UpdateProfile))).Methods("PUT")
	twwetRouter := router.PathPrefix("tweet").Subrouter()
	twwetRouter.HandleFunc("/", middleware.CheckDB(middleware.ValidJWT(routers.CreateTweet))).Methods("POST")
	twwetRouter.HandleFunc("/", middleware.CheckDB(middleware.ValidJWT(routers.GetTweets))).Methods("GET")
	twwetRouter.HandleFunc("/", middleware.CheckDB(middleware.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
