package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomePageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the library API!"}`
	r := SetUpRouter()
	r.GET("/", HomePageHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test other functions
// See "comprehensive guide to testing in Go by Sergey & Alexandre" to try table-driven test and package test
// See "Testcontainers for Go" to be more systematic about testing
