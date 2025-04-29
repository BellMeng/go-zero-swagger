package goZeroSwagger

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag"
)

type mockedSwag struct{}

func (s *mockedSwag) ReadDoc() string {
	return `{}`
}

func TestWrapHandler(t *testing.T) {
	var restConf rest.RestConf
	restConf.Port = 8888
	s, err := rest.NewServer(restConf)
	if err != nil {
		assert.NotNil(t, err)
		return
	}

	s.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/:any",
			Handler: WrapHandler(swaggerFiles.Handler, URL("https://github.com/swaggo/gin-swagger")),
		},
	)

	defer s.Stop()
	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	resp := httptest.NewRecorder()

	s.Routes()[0].Handler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestWrapCustomHandler(t *testing.T) {
	var restConf rest.RestConf
	restConf.Port = 8888
	s, err := rest.NewServer(restConf)
	if err != nil {
		assert.NotNil(t, err)
		return
	}

	s.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/:any",
			Handler: WrapHandler(swaggerFiles.Handler, URL("https://github.com/swaggo/gin-swagger")),
		},
	)

	defer s.Stop()

	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	resp := httptest.NewRecorder()

	s.Routes()[0].Handler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "text/html; charset=utf-8")

	req = httptest.NewRequest(http.MethodGet, "/doc.json", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	doc := &mockedSwag{}
	swag.Register(swag.Name, doc)

	req = httptest.NewRequest(http.MethodGet, "/doc.json", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "application/json; charset=utf-8")

	// Perform body rendering validation
	w2Body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, doc.ReadDoc(), string(w2Body))

	req = httptest.NewRequest(http.MethodGet, "/favicon-16x16.png", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "image/png")

	req = httptest.NewRequest(http.MethodGet, "/swagger-ui.css", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "text/css; charset=utf-8")

	req = httptest.NewRequest(http.MethodGet, "/swagger-ui-bundle.js", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "application/javascript")

	req = httptest.NewRequest(http.MethodGet, "/index.css", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "text/css; charset=utf-8")

	req = httptest.NewRequest(http.MethodGet, "/swagger-initializer.js", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, resp.Header()["Content-Type"][0], "application/javascript")

	req = httptest.NewRequest(http.MethodGet, "/notfound", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	req = httptest.NewRequest(http.MethodPost, "/index.html", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	req = httptest.NewRequest(http.MethodPut, "/index.html", nil)
	resp = httptest.NewRecorder()
	s.Routes()[0].Handler(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
