package api

import (
	"go-microservice/user-service/api/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	rootRouter := chi.NewRouter()
	userRouter := chi.NewRouter()

	rootRouter.Mount("/api/v1", userRouter)

	userRouter.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetListUserHandler)
		r.Post("/", handler.CreateUserHandler)
	})

	return rootRouter
}
