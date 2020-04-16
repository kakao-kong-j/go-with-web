package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestBarPathHandler_withOutJSON(t *testing.T) {
	assert := assert.New(t)

	res:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/foo", nil)

	mux:=NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusBadRequest,res.Code)
}

func TestBarPathHandler_withJSON(t *testing.T) {
	assert := assert.New(t)

	res:=httptest.NewRecorder()
	req:=httptest.NewRequest("GET","/foo", strings.NewReader(`{"first_name":"jinho", "last_name":"hong","email":"tpdleps@gmail.com"}`))


	mux:=NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusCreated,res.Code)

	user :=new(User)
	err:=json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("jinho",user.FirstName)
	assert.Equal("hong",user.LastName)
}