package util_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	boot "github.com/summerKK/mall-api/init"
	"github.com/summerKK/mall-api/internal/middleware"
	"github.com/summerKK/mall-api/pkg/util"
)

func init() {
	boot.SetConfig([]string{"../../configs"})
	boot.Boot()
}

func TestAddErrorToCtx(t *testing.T) {
	engine := gin.Default()
	r := engine.Use(middleware.CollectError(nil))

	r.GET("/test0", func(c *gin.Context) {
		var err error
		defer func() {
			util.AddErrorToCtx(c, err)
		}()

		err = errors.New("test")
	})

	req0 := httptest.NewRequest("GET", "/test0", nil)
	w0 := httptest.NewRecorder()
	engine.ServeHTTP(w0, req0)
	assert.Equal(t, w0.Code, http.StatusOK)

	r.GET("/test1", func(c *gin.Context) {
		util.AddErrorToCtx(c, errors.New(""))
	})

	req1 := httptest.NewRequest("GET", "/test1", nil)
	w1 := httptest.NewRecorder()
	engine.ServeHTTP(w1, req1)
	assert.Equal(t, w1.Code, http.StatusOK)
}
