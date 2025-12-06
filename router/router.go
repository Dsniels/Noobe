package router

import (
	"net/http"

	"github.com/dsniels/noobe/handlers"
	mdw "github.com/dsniels/noobe/middleware"
)

func InitRoutes() http.Handler {

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HomepageHandler)

	return mdw.ExceptionMiddleware(router)
}
