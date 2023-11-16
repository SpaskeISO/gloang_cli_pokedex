package pokeapi

import (
	"net/http"
	"time"

	"github.com/SpaskeISO/golang_cli_pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cashe      pokecache.Cashe
	httpClient http.Client
}

func NewClient(casheInterval time.Duration) Client {
	return Client{
		cashe: pokecache.NewCashe(casheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
