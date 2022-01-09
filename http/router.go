package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zknill/vending/domain"
)

// NOTE: This is a **toy** server implementation.
// The underlying machine holds state in memory, and
// isn't threadsafe to be used in an http server.
// This server is here to show how you _could_ interact
// with the machine; including endpoints, request and
// response bodies, and status codes.
// The implementation is not complete.

type Router struct {
	machine domain.Machine
}

func NewRouter(machine domain.Machine) Router {
	return Router{machine: machine}
}

func (r Router) Mount() *chi.Mux {
	server := chi.NewRouter()

	server.Get("/products", products(r.machine))
	server.Post("/purchase", purchase(r.machine))

	return server
}

type productItem struct {
	Coordinate string `json:"coordinate"`
	Price      string `json:"price"`
	Name       string `json:"name"`
}

func products(machine domain.Machine) http.HandlerFunc {
	type response struct {
		Products []productItem `json:"products"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		p := machine.Products()

		out := make([]productItem, len(p))

		for i := range p {
			pp := p[i]

			out[i] = productItem{
				Coordinate: pp.Coordinate,
				Price:      formatPrice(pp.Price),
				Name:       pp.Name,
			}
		}

		writeJson(w, response{Products: out})
	}
}

func formatPrice(coin int) string {
	return fmt.Sprintf("$%.2f", float64(coin)/100)
}

type errorMessage struct {
	Message       string `json:"message"`
	MoneyRequired int    `json:"money_required,omitempty"`
}

func purchase(machine domain.Machine) http.HandlerFunc {
	type request struct {
		Coins      []int  `json:"coins"`
		Coordinate string `json:"coordinate"`
	}

	type response struct {
		Product productItem `json:"product"`
		Change  []uint      `json:"change"`
	}

	// NOT thread safe!
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, coin := range req.Coins {
			if coin < 0 {
				w.WriteHeader(http.StatusBadRequest)
				writeJson(w, errorMessage{
					Message: fmt.Sprintf("negative coin: %d", coin),
				})
				return
			}

			if err := machine.InsertCoin(uint(coin)); err != nil {
				w.WriteHeader(http.StatusBadRequest)

				var unknown domain.ErrUnknownCoin
				if errors.As(err, &unknown) {
					writeJson(w, errorMessage{
						Message: fmt.Sprintf("unknown coin: %d", unknown.UnknownCoin()),
					})
					return
				}
			}
		}

		product, change, err := machine.Purchase(req.Coordinate)
		if err != nil {
			handlePurchaseError(w, err)
			return
		}

		writeJson(w, response{
			Product: productItem{
				Coordinate: product.Coordinate,
				Price:      formatPrice(product.Price),
				Name:       product.Name,
			},
			Change: change,
		})
	}
}

func handlePurchaseError(w http.ResponseWriter, err error) {
	var (
		stock   domain.ErrOutOfStock
		money   domain.ErrNotEnoughMoney
		product domain.ErrUnknownProduct
	)

	switch {
	case errors.As(err, &stock):
		w.WriteHeader(http.StatusGone)

		writeJson(w, errorMessage{
			Message: fmt.Sprintf("product out of stock: %q", stock.OutOfStockProduct()),
		})

		return
	case errors.As(err, &money):
		w.WriteHeader(http.StatusPaymentRequired)

		writeJson(w, errorMessage{
			Message: fmt.Sprintf("not enough money, more required: $%.2f",
				float64(money.RemainingRequired())/100),
			MoneyRequired: money.RemainingRequired(),
		})

		return
	case errors.As(err, &product):
		w.WriteHeader(http.StatusNotFound)

		writeJson(w, errorMessage{
			Message: fmt.Sprintf("product not found: %q", product.UnknownProduct()),
		})

		return

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func writeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
