package signal

import (
	"github.com/gin-gonic/gin"
	"gogc/src/model"
)

func RequestInit(c *gin.Context) model.Request {

	request := model.Request{}
	request.Params = map[string]string{}
	request.Query = c.Request.URL.Query()

	return request
}
