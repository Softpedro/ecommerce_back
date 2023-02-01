package paypal

import (
	"io"
	"log"
	"net/http"

	"github.com/Softpedro/ecommerce_back/domain/paypal"
	"github.com/Softpedro/ecommerce_back/infrastructure/handler/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCasePayPal paypal.UseCase
	responser     response.API
}

func newHandler(ucp paypal.UseCase) handler {
	return handler{useCasePayPal: ucp}
}

func (h handler) Webhook(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return h.responser.BindFailed(err)
	}
	go func() {
		err = h.useCasePayPal.ProcessRequest(c.Request().Header, body)
		if err != nil {
			log.Printf("useCasePayPal.ProcessRequest(): %v", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
