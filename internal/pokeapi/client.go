package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokecache"
)

// Client handles HTTP requests to the PokéAPI with caching functionality.
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

/*
NewClient creates and returns a new Client instance.

The client includes a cache for API responses and a configured HTTP client.

Parameters:
- timeout: The maximum duration for an HTTP request.
- cacheInterval: The expiration interval for cached responses.

Returns:
- Client: A new API client instance.
*/
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

/*
ListLocations retrieves a paginated list of location areas from the PokéAPI.

If the requested page URL is cached, the cached response is returned.

Parameters:
- pageURL: An optional pointer to a string representing the next/previous page URL.

Returns:
- LocationAreaList: The response containing location areas.
- error: An error if the request or JSON decoding fails.
*/
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	locationsResp := LocationAreaList{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(url, body)
	return locationsResp, nil
}

/*
GetLocation retrieves detailed information about a specific location area.

Parameters:
- locationName: The name of the location area to fetch.

Returns:
- LocationArea: The response containing location details.
- error: An error if the request or JSON decoding fails.
*/
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResp := LocationArea{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, body)

	return locationResp, nil
}

/*
GetPokemon retrieves details about a specific Pokémon by name.

Parameters:
- pokemonName: The name of the Pokémon to fetch.

Returns:
- Pokemon: The response containing Pokémon details.
- error: An error if the request or JSON decoding fails.
*/
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemonResp, nil
}
