package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) (*Client, error) {
	cache, err := pokecache.NewCache(5 * time.Second)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}, nil
}
