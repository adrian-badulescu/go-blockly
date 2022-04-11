package blockly

import (
	"fmt"
	"net/http"
	"os"

	// "path/filepath"
	"github.com/gin-gonic/gin"
)

func RegisterBlocklyAutomationEndpoint(c *gin.Engine) {
	// router.Load
	// c.HTML(http.StatusOK, "index.html", nil)
	c.StaticFS("/BlocklyAutomation", http.Dir(os.Getenv("GOPATH")))

	c.StaticFS("/BlocklyAutomationa", http.Dir("."))
	c.StaticFS("/BlocklyAutomationb", http.Dir(os.Getenv("GOPATH")+"/pkg/mod/github.com/adrian-badulescu/go-blockly@v0.0.0-20220411174346-b8e69f92bb55"))

	fmt.Println(http.Dir(os.Getenv("GOPATH")))

	// fmt.Println("a:", (http.Dir(filepath.Dir(os.Getenv("GOPATH")))), filepath.Dir("/pkg/mod/github.com/adrian-badulescu/go-blockly@v0.0.0-20220313143102-dabce715a36a"))
}

// func main() {
// 	router := gin.Default()
// 	RegisterBlocklyAutomationEndpoint(router)
// }
