package paypal

import (
	"net/http"

	"github.com/Softpedro/ecommerce_back/model"
	"github.com/google/uuid"
)

type UseCase interface {
	ProcessRequest(header http.Header, body []byte) error
}

type UseCasePurchaseOrder interface {
	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}

type UseCaseInvoice interface {
	Create(m *model.PurchaseOrder) error
}
