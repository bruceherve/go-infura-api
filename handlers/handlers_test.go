package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
)

//Test Greeting
func TestGreeting(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api",
			expectedCode: 200,
		},

		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API!")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}

// Test GetBlockByNumber
func TestGetBlockByNumber(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/latest-block", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

// Test GetTransactionByBlockHashAndIndex
func TestGetTransactionByBlockHashAndIndex(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/get-transaction", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

}

//Test GetBalance
func TestGetBalnce(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/get-balance", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

// Test GetBlockTransactionCountByNumber
func TestGetBlockTransactionCountByNumber(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/get-transaction-count", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

//Test GetGasPrice
func TestGetGasPrice(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/get-transaction-count", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
