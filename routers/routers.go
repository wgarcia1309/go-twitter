package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wgarcia1309/go-twitter/handlers"
	"github.com/wgarcia1309/go-twitter/middleware"
)

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/login", middleware.CheckDB(handlers.Login)).Methods("POST")

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("", handlers.NewUser).Methods("POST")
	userRouter.HandleFunc("", middleware.ValidJWT(handlers.GetProfile)).Methods("GET")
	userRouter.HandleFunc("", middleware.ValidJWT(handlers.UpdateProfile)).Methods("PUT")

	twwetRouter := router.PathPrefix("/tweet").Subrouter()
	twwetRouter.HandleFunc("", middleware.CheckDB(middleware.ValidJWT(handlers.CreateTweet))).Methods("POST")
	twwetRouter.HandleFunc("", middleware.CheckDB(middleware.ValidJWT(handlers.GetTweets))).Methods("GET")
	twwetRouter.HandleFunc("", middleware.CheckDB(middleware.ValidJWT(handlers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/follow", middleware.CheckDB(middleware.ValidJWT(handlers.Follow))).Methods("POST")
	router.HandleFunc("/unfollow", middleware.CheckDB(middleware.ValidJWT(handlers.Unfollow))).Methods("DELETE")
	router.HandleFunc("/following", middleware.CheckDB(middleware.ValidJWT(handlers.GetRelation))).Methods("GET")

	router.HandleFunc("/search", middleware.CheckDB(middleware.ValidJWT(handlers.FindUsers))).Methods("GET")
	router.HandleFunc("/feed", middleware.CheckDB(middleware.ValidJWT(handlers.GetFeed))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
