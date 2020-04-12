package myapp

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/", nil)

	mux:=NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusOK,res.Code)
	data,_:=ioutil.ReadAll(res.Body)
	assert.Equal("Hello World",string(data))
}

func TestBarPathHandler(t *testing.T) {
	assert := assert.New(t)

	res:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/bar", nil)

	mux:=NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusOK,res.Code)
	data,_:=ioutil.ReadAll(res.Body)
	assert.Equal("Hello world!",string(data))
}

func TestBarPathHandler_withOutName(t *testing.T) {
	assert := assert.New(t)

	res:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/bar?name=jinho", nil)

	mux:=NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusOK,res.Code)
	data,_:=ioutil.ReadAll(res.Body)
	assert.Equal("Hello jinho!",string(data))
}
