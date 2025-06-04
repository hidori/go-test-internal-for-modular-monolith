package main

import (
	router1 "github.co/hidori/go-test-internal-for-modular-monolith/api1/infra/router"
	router2 "github.co/hidori/go-test-internal-for-modular-monolith/api2/infra/router"
	"github.co/hidori/go-test-internal-for-modular-monolith/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.Public()
	//internal1.Public() // not allowed

	server := gin.Default()

	rg1 := server.Group("/app1")
	router1.Attach(rg1)

	rg2 := server.Group("/app2")
	router2.Attach(rg2)

	server.Run(":8080")
}
