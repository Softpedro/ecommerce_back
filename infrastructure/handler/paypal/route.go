package paypal

import (
	"github.com/Softpedro/ecommerce_back/domain/invoice"
	"github.com/Softpedro/ecommerce_back/domain/paypal"
	"github.com/Softpedro/ecommerce_back/domain/purchaseorder"
	storageInvoice "github.com/Softpedro/ecommerce_back/infrastructure/postgres/invoice"
	storagePurchaseOrder "github.com/Softpedro/ecommerce_back/infrastructure/postgres/purchaseorder"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	purchaseOrderUseCase := purchaseorder.New(storagePurchaseOrder.New(dbPool))
	invoiceUseCase := invoice.New(storageInvoice.New(dbPool), nil)
	useCase := paypal.New(purchaseOrderUseCase, invoiceUseCase)

	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/paypal")

	route.POST("", h.Webhook)
}
