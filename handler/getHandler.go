package handler

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type record struct {
	nTimes         int
	expirationTime time.Time
}

var lock sync.Mutex
var m map[string]record = map[string]record{}

// GetHandler handle / request
func (h Handler) GetHandler(c *fiber.Ctx) error {
	ip := c.IP()

	lock.Lock()
	r := m[ip]

	if time.Now().After(r.expirationTime) {
		r.nTimes = 1
		r.expirationTime = time.Now().Add(time.Duration(h.Config.ExpirationSecond) * time.Second)
	} else {
		r.nTimes++
	}

	m[ip] = r

	lock.Unlock()

	if r.nTimes > h.Config.LimitTimes {
		return c.SendString("Error")
	}

	return c.SendString(strconv.Itoa(r.nTimes))
}
