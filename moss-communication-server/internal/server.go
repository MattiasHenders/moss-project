package internal

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/MattiasHenders/moss-communication-server/internal/handlers"
	h "github.com/MattiasHenders/moss-communication-server/pkg/handler"
	"github.com/MattiasHenders/moss-communication-server/pkg/middleware"
	pkgMiddleware "github.com/MattiasHenders/moss-communication-server/pkg/middleware"
	"github.com/MattiasHenders/moss-communication-server/pkg/secrets"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Start(port string) {

	secrets := secrets.LoadEnvAndGetSecrets()

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(pkgMiddleware.Logger("/"))
	r.Use(chiMiddleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// Authenticated Api Key routes here...
	r.Group(func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.VerifyAPIKey(secrets.DemoAPIKey))

			// Stable Diffusion routes...
			r.Post("/txt-2-img", h.Handler(handlers.CreateTextToImageRequestHandler()))
			r.Post("/img-2-img", h.Handler(handlers.CreateImageToImageRequestHandler()))
		})
	})

	// Health check routes here...
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	zap.L().Info(fmt.Sprintf("Starting moss-communication-server server on port %s...", port))
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
