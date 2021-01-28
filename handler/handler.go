package handler

// Config config handler
type Config struct {
	LimitTimes       int
	ExpirationSecond int
}

// Handler instance
type Handler struct {
	Config *Config
}
