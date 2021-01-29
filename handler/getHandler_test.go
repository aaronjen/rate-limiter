package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// go test -run Test_Get_Handler -race -v ./handler
func Test_Get_Handler(t *testing.T) {
	app := fiber.New()

	h := Handler{
		&Config{
			LimitTimes:       10,
			ExpirationSecond: 3,
		},
	}

	app.Get("/", h.GetHandler)

	wg := sync.WaitGroup{}

	reqFunc := func() {
		defer wg.Done()
		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		assert.Equal(t, nil, err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NotEqual(t, "Error", string(body))
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reqFunc()
	}

	wg.Wait()

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Equal(t, nil, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "Error", string(body))

	time.Sleep(5 * time.Second)
	resp, err = app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Equal(t, nil, err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	assert.NotEqual(t, "Error", string(body))
}
