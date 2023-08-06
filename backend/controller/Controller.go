package controller

import (
	"backend/algorithm"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller(c *gin.Context) {
	var input models.InputGraph
	var tarjans algorithm.Tarjans

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.OutputGraph{
			Graph:  nil,
			Status: err.Error(),
		})
		return
	}

	value := utils.ChangeToMultipleChild(input.Input)

	for _, val := range value {
		tarjans.Push(val)
	}

	if input.Type == "bridge" {
		val, _ := tarjans.Bridge()
		c.JSON(http.StatusOK, models.OutputGraph{
			Graph:  val,
			Status: "ok",
		})
		return
	}
	if input.Type == "scc" {
		val, _ := tarjans.SCC()
		c.JSON(http.StatusOK, models.OutputGraph{
			Graph:  val,
			Status: "ok",
		})
		return
	}

	c.JSON(http.StatusBadRequest, models.OutputGraph{
		Graph:  nil,
		Status: "no output",
	})
	return
}
