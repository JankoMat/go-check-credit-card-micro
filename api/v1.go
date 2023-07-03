package api

import (
	"fmt"
	"net/http"

	"github.com/JankoMat/go-check-credit-card-micro/pkg"
	"github.com/gin-gonic/gin"
)

func path(endpoint string) string {
	return fmt.Sprintf("/api/v1/%s", endpoint)
}

type RESTApiV1 struct {
	router *gin.Engine
}

func (api *RESTApiV1) Serve(addr string) error {
	return api.router.Run(addr)
}

func NewRESTApiV1() *RESTApiV1 {
	router := gin.Default()
	api := &RESTApiV1{
		router,
	}

	router.POST(path("check/:card"), api.CheckCard)
	router.GET(path("healthz"), api.Health)

	return api
}

func (api *RESTApiV1) CheckCard(c *gin.Context) {
	card := c.Param("card")

	isValid := card.Valid(card)
	c.JSON(http.StatusOK, gin.H{
		"cardIsValid": isValid,
	})
}

func (api *RESTApiV1) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Healthy",
	})
}
