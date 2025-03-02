package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2/location-area"
	limit   = 20
)

var (
	offset      = 0
	isFirstPage = true
	firstPage   = fmt.Sprintf("%s?offset=%d&limit=%d", baseURL, offset, limit)
)

func fetchLocationAreas(cfg *config, cache *pokecache.Cache, apiURL string) error {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, cfg); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	cache.Add(apiURL, body)

	return nil
}

func fetchLocationAreasFromCache(cfg *config, cache *pokecache.Cache, apiURL string) bool {
	if val, exists := cache.Get(apiURL); exists {
		if err := json.Unmarshal(val, cfg); err != nil {
			return false
		}
		return true
	}

	return false
}

func commandMap(cfg *config, cache *pokecache.Cache) error {
	isFirstPage = false
	url := cfg.Next
	if cfg.Next == "" && cfg.Previous == "" {
		url = fmt.Sprintf("%s?offset=%d&limit=%d", baseURL, offset, limit)
	} else if cfg.Next == "" {
		fmt.Println("No more location areas to show.")
		return nil
	}

	if ok := fetchLocationAreasFromCache(cfg, cache, url); !ok {
		err := fetchLocationAreas(cfg, cache, url)
		if err != nil {
			return err
		}
	}

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	if isFirstPage || cfg.Previous == "" {
		return errors.New("you're on the first page")
	}

	if cfg.Previous == firstPage {
		isFirstPage = true
	}

	if ok := fetchLocationAreasFromCache(cfg, cache, cfg.Previous); !ok {
		err := fetchLocationAreas(cfg, cache, cfg.Previous)
		if err != nil {
			return err
		}
	}

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}

	if isFirstPage {
		cfg.Previous = ""
	}

	return nil
}
