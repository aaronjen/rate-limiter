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

var m sync.Map = sync.Map{}

// GetHandler handle / request
func (h Handler) GetHandler(c *fiber.Ctx) error {
	ip := c.IP()

	raw, _ := m.LoadOrStore(ip, record{
		nTimes:         0,
		expirationTime: time.Now().Add(time.Duration(h.Config.ExpirationSecond) * time.Second),
	})

	r := raw.(record)

	if time.Now().After(r.expirationTime) {
		r.nTimes = 1
		r.expirationTime = time.Now().Add(time.Duration(h.Config.ExpirationSecond) * time.Second)
	} else {
		r.nTimes++
	}

	m.Store(ip, r)

	if r.nTimes > h.Config.LimitTimes {
		return c.SendString("Error")
	}

	return c.SendString(strconv.Itoa(r.nTimes))
}
