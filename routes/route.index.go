package routes

import (
	"fmt"
	"log"
	"net/http"
	"sv-article/pkg/configs"

	"github.com/rs/cors"
)

func LaunchApp() {
	WebServer := fmt.Sprintf("Wev Server is Running on Port %v", configs.WEB_PORT)
	log.Println(WebServer)
	handler := cors.AllowAll().Handler(Route())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configs.WEB_PORT), handler))
}