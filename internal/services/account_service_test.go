package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
    req, _ := http.NewRequest("GET", "/balance/1", nil)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = req

    GetBalance(c)

    assert.Equal(t, http.StatusOK, w.Code)
}
