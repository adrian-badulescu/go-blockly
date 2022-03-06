package blockly

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBlocklyAutomationEndpoint(c *gin.Engine) {
	// router.Load
	// c.HTML(http.StatusOK, "index.html", nil)
	c.StaticFS("/BlocklyAutomation", http.Dir("./blockly"))
}
