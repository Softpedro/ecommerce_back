package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/Softpedro/ecommerce_back/domain/user"
	"github.com/Softpedro/ecommerce_back/infrastructure/handler/middle"
	storageUser "github.com/Softpedro/ecommerce_back/infrastructure/postgres/user"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/admin/users", middlewares...)

	g.GET("", h.GetAll)

	i := e.Group("/api/v1/private/user", middlewares...)

	i.GET("", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")

	g.POST("", h.Create)
}
