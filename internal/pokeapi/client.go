package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaList{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	locationsResp := LocationAreaList{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

// GetLocation -
func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResp := LocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
