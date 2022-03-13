package blockly

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func RegisterBlocklyAutomationEndpoint(c *gin.Engine) {
	// router.Load
	// c.HTML(http.StatusOK, "index.html", nil)
	c.StaticFS("/BlocklyAutomation", http.Dir(os.Getenv("GOPATH")))
}
