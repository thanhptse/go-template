package contxt

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type contextKey int

var appContextKey contextKey

type AppContext struct {
	ginContext    *gin.Context
	prefix        string
	requestInfo   RequestInfo
	parentContext context.Context
}

type RequestInfo struct {
	RequestID string
	ClientIP  string
	Host      string
	Method    string
	Path      string
	Referer   string
	UserAgent string
}

func NewAppContext(c context.Context) IAppContext {
	ginContext, ok := c.(*gin.Context)
	if ok {
		c = ginContext.Request.Context()
	}
	value := c.Value(appContextKey)
	if value == nil {
		return &AppContext{requestInfo: RequestInfo{RequestID: uuid.NewString()}}
	}

	ctx, ok := value.(*AppContext)
	if !ok {
		return &AppContext{requestInfo: RequestInfo{RequestID: uuid.NewString()}}
	}

	ctx.parentContext = c
	return ctx
}

// SetupAppContext is a middleware to embbed this Context type into gin.Context
func SetupAppContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), appContextKey, initContext(c))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func initContext(c *gin.Context) *AppContext {
	var requestID string
	if reqID, ok := c.Get("RequestID"); ok {
		requestID, _ = reqID.(string)
	}

	ctx := &AppContext{
		ginContext: c,
	}
	if c.Request != nil {
		ctx.requestInfo = RequestInfo{
			RequestID: requestID,
			ClientIP:  c.ClientIP(),
			Host:      c.Request.Host,
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Referer:   c.Request.Referer(),
			UserAgent: c.Request.UserAgent(),
		}
		ctx.parentContext = c.Request.Context()
	}

	return ctx
}
