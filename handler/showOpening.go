package handler

import (
	"net/http"

	"github.com/gaspartv/API-GO-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context){
	id := ctx.Query("id")
	if id == "" {
			sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
			return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, "opening not found")
        return
    }

    sendSuccess(ctx, "show opening", opening)
}
