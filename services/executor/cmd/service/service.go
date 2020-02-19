package cmd

import (
	"fmt"

	"gitlab.com/scalent/goxpert/services/executor/pkg/v1/endpoints"
	"gitlab.com/scalent/goxpert/services/executor/pkg/v1/middleware"

	handler "gitlab.com/scalent/goxpert/services/executor/pkg/v1/handlers"

	service "gitlab.com/scalent/goxpert/services/executor/pkg/v1/services"

	repository "gitlab.com/scalent/goxpert/services/executor/pkg/v1/repositories"

	"net/http"

	"github.com/gorilla/mux"
)

//Run **
func Run() {
	router := mux.NewRouter()

	//User Services dependancies
	executorRepository := repository.NewExecutorRepositoryImpl()
	executorService := service.NewExecutorServiceImpl(executorRepository)
	executorHandler := handler.NewExecutorHandlerImpl(executorService)

	endpoints.NewExecutorRoute(router, executorHandler)

	router.Use(middleware.LoggingMiddleware)
	//router.Use(AuthMiddleware)
	fmt.Println("Service Running!")
	fmt.Println(http.ListenAndServe(":9001", router))
}
