package apis

import (
	"go-admin/tools/app"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	target := "9.135.8.235:8080"
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func GetArticleList(c *gin.Context) {
	var res app.Response
	res.Data = "hello world"

	c.JSON(http.StatusOK, res.ReturnOK())
}
