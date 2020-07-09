package middlewares

import (
	"net/http"
	test "net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func CORSTest(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	res := test.NewRecorder()

	c := e.NewContext(req, res)
	m := InitMiddleware()

	h := m.CORS(echo.HandlerFunc(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}))

	err := h(c)

	if err != nil {
		t.Fatal(err)
	}
}
