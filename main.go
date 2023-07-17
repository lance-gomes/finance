package main

import (
	"finance/config"
	"finance/plaid"
	"finance/server"
)

func main() {
	conf := config.New()
	plaid := plaid.Init(*conf)

	server.Serve(conf, plaid)
}
