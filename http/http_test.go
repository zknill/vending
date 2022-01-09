package http_test

import (
	"io/ioutil"
	nethttp "net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/vending/domain"
	"github.com/zknill/vending/domain/domainfakes"
	"github.com/zknill/vending/http"
)

var _ = Describe("Http", func() {
	var (
		router  http.Router
		machine *domainfakes.FakeMachine
	)

	BeforeEach(func() {
		machine = new(domainfakes.FakeMachine)
		router = http.NewRouter(machine)
	})

	Describe("Products", func() {
		It("fetches products", func() {
			machine.ProductsReturns([]domain.Product{
				{
					Coordinate: "A1",
					Price:      125,
					Name:       "Fruit Pastels",
				},
				{
					Coordinate: "A2",
					Price:      90,
					Name:       "Walkers Crisps",
				},
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/products", nil)

			router.Mount().ServeHTTP(w, r)
			b, _ := ioutil.ReadAll(w.Result().Body)

			expected := `
{
  "products": [
    {
      "coordinate": "A1",
      "price": "$1.25",
      "name": "Fruit Pastels"
    },
    {
      "coordinate": "A2",
      "price": "$0.90",
      "name": "Walkers Crisps"
    }
  ]
}`

			Expect(w.Result().StatusCode).To(Equal(200))
			Expect(b).To(MatchJSON(expected))
		})
	})

	Describe("Purchase", func() {
		var (
			w *httptest.ResponseRecorder
			r *nethttp.Request
			p = domain.Product{
				Coordinate: "A2",
				Price:      90,
				Name:       "Walkers Crisps",
			}
		)
		BeforeEach(func() {
			w = httptest.NewRecorder()
			machine.PurchaseReturns(p, []uint{25, 50, 5}, nil)
			body := `{
    "coins": [100],
    "coordinate": "A2"
}`
			r = httptest.NewRequest("POST", "/purchase", strings.NewReader(body))
		})

		It("Purchases a product and receives change", func() {
			router.Mount().ServeHTTP(w, r)
			Expect(w.Result().StatusCode).To(Equal(200))

			expected := `
{
  "product": {
      "coordinate": "A2",
      "price": "$0.90",
      "name": "Walkers Crisps"
  },
  "change": [25, 50, 5]
}`
			b, _ := ioutil.ReadAll(w.Result().Body)
			Expect(b).To(MatchJSON(expected))
		})

		It("Handles not enough money", func() {
			machine.PurchaseReturns(domain.Product{}, nil, errMoney(23))
			router.Mount().ServeHTTP(w, r)
			Expect(w.Result().StatusCode).To(Equal(402))

			expected := `{"message":"not enough money, more required: $0.23","money_required":23}`
			b, _ := ioutil.ReadAll(w.Result().Body)
			Expect(b).To(MatchJSON(expected))
		})

		It("Handles not enough stock", func() {
			machine.PurchaseReturns(domain.Product{}, nil, errStock("A2"))
			router.Mount().ServeHTTP(w, r)
			Expect(w.Result().StatusCode).To(Equal(410))

			expected := `{"message":"product out of stock: \"A2\""}`
			b, _ := ioutil.ReadAll(w.Result().Body)
			Expect(b).To(MatchJSON(expected))
		})

		It("Handles unknown stock", func() {
			machine.PurchaseReturns(domain.Product{}, nil, errUnknown("A2"))
			router.Mount().ServeHTTP(w, r)
			Expect(w.Result().StatusCode).To(Equal(404))

			expected := `{"message":"product not found: \"A2\""}`
			b, _ := ioutil.ReadAll(w.Result().Body)
			Expect(b).To(MatchJSON(expected))
		})
	})
})

type errMoney int

func (e errMoney) Error() string          { return "not enough money" }
func (e errMoney) RemainingRequired() int { return int(e) }

type errStock string

func (e errStock) Error() string             { return "out of stock" }
func (e errStock) OutOfStockProduct() string { return string(e) }

type errUnknown string

func (e errUnknown) Error() string          { return "unknown product" }
func (e errUnknown) UnknownProduct() string { return string(e) }
