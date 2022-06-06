package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET
//Path to the root path.
func Greeting(c *fiber.Ctx) error {
	return c.SendString("Welcome to the jsonrpc API!")
}

// swagger:route POST
// Function to get latest block.
//By block number
func GetBlockByNumber(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	url := "https://mainnet.infura.io/v3"
	method := "POST"

	payload := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_getBlockByNumber",
		"params":[
			"0x1b4",
			false
		],
		"id":1
	}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	return ctx.Status(res.StatusCode).SendStream(res.Body)

}

// swagger:route POST
//Function to get transaction.
//By BlockHashAndIndex

func GetTransactionByBlockHashAndIndex(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	url := "https://mainnet.infura.io/v3/"
	method := "POST"

	payload := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_getTransactionByBlockHashAndIndex",
		"params":[
			"0x3c82bc62179602b67318c013c10f99011037c49cba84e31ffe6e465a21c521a7",
			"0x0"
		],
		"id":1
	}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(res.StatusCode).SendStream(res.Body)
}

// swagger:route POST
// Function to GetBalance.
func GetBalance(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	url := "https://mainnet.infura.io/v3/"
	method := "POST"

	payload := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_getBalance",
		"params":[
			"0xc94770007dda54cF92009BFF0dE90c06F603a09f",
			"latest"
		],
		"id":1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(res.StatusCode).SendStream(res.Body)

}

// swagger:route POST
// Function to GetBlockTransactionByCountNumber.

func GetBlockTransactionCountByNumber(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	url := "https://mainnet.infura.io/v3/"
	method := "POST"

	payload := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_getBlockTransactionCountByNumber",
		"params":[
			"latest"
		],
		"id":1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(res.StatusCode).SendStream(res.Body)
}

// swagger:route POST
// Function to GetGasPrice.
func GetGasPrice(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	url := "https://mainnet.infura.io/v3/"
	method := "POST"

	payload := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_gasPrice",
		"params":[],
		"id":1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(res.StatusCode).SendStream(res.Body)

}
