package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/udm/internal/logger"
)

// Unsubscribe - unsubscribe from notifications
func (p *Processor) HandleUnsubscribe(c *gin.Context) {
	logger.SdmLog.Infof("Handle Unsubscribe")

	// step 2: retrieve request
	supi := c.Params.ByName("supi")
	subscriptionID := c.Params.ByName("subscriptionId")

	// step 3: handle the message
	problemDetails := p.consumer.UnsubscribeProcedure(supi, subscriptionID)

	// step 4: process the return value from step 3
	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	} else {
		c.Status(http.StatusNoContent)
		return
	}
}
