package main

import (
	"backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	if err := r.Run(); err != nil {
		return
	}
}
