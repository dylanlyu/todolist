package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func SetGinRequestData(gc *gin.Context,data []byte) *gin.Context {
	gc.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	gc.Next()
	return gc
}