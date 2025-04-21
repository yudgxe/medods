package router

import (
	"medods/http/handlers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func New() *chi.Mux {
	root := chi.NewRouter()
	{
		root.Use(middleware.Recoverer)
		root.Use(middleware.RequestID)
		root.Use(middleware.RedirectSlashes)
		root.Use(render.SetContentType(render.ContentTypeJSON))
		root.Use(middleware.Timeout(time.Minute))
	}

	v1 := chi.NewRouter()

	handlers.BindAuth(v1)

	root.Mount("/api/v1", v1)

	return root
}
