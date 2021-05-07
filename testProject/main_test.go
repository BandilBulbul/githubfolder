package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The total  should not be  nil")
	assert.Equal(t, 4, total, "expecting 4")

}
func TestSubtract(t *testing.T) {
	total := Subtract(10, 3)
	assert.NotNil(t, total, "total should not be nil")
	assert.Equal(t, 7, total, "expecting 7")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}
func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok response is expected")
	assert.Equal(t, "Hello World", response.Body.String(), "incorrect body found")
	//fmt.Println(response.Body)

}

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2+2 to equal 4")
	}
}

func TestCalculateTestify(t *testing.T) {
	assert.NotEqual(t, Calculate(4), "down")
	assert.Equal(t, Calculate(4), 6)

}

//Table Driven testing
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}

func TestCalculateWithTestify(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
