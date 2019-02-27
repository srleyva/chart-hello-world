package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMachine(t *testing.T) {
	machine := Vending{
		coins:     0,
		chocolate: 10,
		cost:      1,
	}
	e := echo.New()
	t.Run("Test machine Get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		machine.Get(c)
		if rec.Code != 200 {
			t.Errorf("non 200 returned where expected: %d", rec.Code)
		}
		expected := "There are 10 pieces of chocolate left."
		actual := rec.Body.String()
		if actual != expected {
			t.Errorf("Expected: %s \n Actual: %s", expected, actual)
		}
	})
	t.Run("Test machine Post", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		machine.Post(c)
		if rec.Code != 200 {
			t.Errorf("non 200 returned where expected: %d", rec.Code)
		}
	})
}
