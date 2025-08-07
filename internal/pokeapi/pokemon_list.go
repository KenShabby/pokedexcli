package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get Pokemon details on an area
func (c *Client) ListPokemon(area string) (RespLocationAreaDetail, error) {
	url := baseURL + "/location-area/"
	if area != "" {
		url += area
	}

	if cached, ok := c.cache.Get(url); ok {
		var pokemonResp RespLocationAreaDetail
		if err := json.Unmarshal(cached, &pokemonResp); err != nil {
			return RespLocationAreaDetail{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreaDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreaDetail{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreaDetail{}, err
	}

	c.cache.Add(url, dat)

	pokemonResp := RespLocationAreaDetail{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespLocationAreaDetail{}, err
	}

	return pokemonResp, nil

}
