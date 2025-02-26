package main

func main() {
	startRepl(&config{})
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	// "net/url"
// 	// "strconv"
// 	"time"
// )

// type LocationArea struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// type Response struct {
// 	Results  []LocationArea `json:"results"`
// 	Next     string         `json:"next"`
// 	Previous string         `json:"previous"`
// }

// const baseURL = "https://pokeapi.co/api/v2/location-area"
// const limit = 20

// var offset = 0 // Tracks the current offset for pagination

// func fetchLocationAreas(client *http.Client, offset int) (Response, error) {
// 	apiURL := fmt.Sprintf("%s?limit=%d&offset=%d", baseURL, limit, offset)
// 	resp, err := client.Get(apiURL)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return Response{}, fmt.Errorf("unexpected status code: %s", resp.Status)
// 	}

// 	var result Response
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return Response{}, err
// 	}

// 	return result, nil
// }

// func nextPage(client *http.Client) {
// 	offset += limit
// 	result, err := fetchLocationAreas(client, offset)
// 	if err != nil {
// 		fmt.Println("Error fetching next page:", err)
// 		return
// 	}

// 	printResults(result)
// }

// func prevPage(client *http.Client) {
// 	if offset-limit < 0 {
// 		fmt.Println("Already at the first page.")
// 		return
// 	}
// 	offset -= limit
// 	result, err := fetchLocationAreas(client, offset)
// 	if err != nil {
// 		fmt.Println("Error fetching previous page:", err)
// 		return
// 	}

// 	printResults(result)
// }

// func printResults(result Response) {
// 	for _, area := range result.Results {
// 		fmt.Println(area.Name, "-", area.URL)
// 	}
// 	fmt.Printf("\nNext: %s\nPrevious: %s\n\n", result.Next, result.Previous)
// }

// func main() {
// 	client := &http.Client{Timeout: 10 * time.Second}

// 	fmt.Println("Fetching initial page...")
// 	result, err := fetchLocationAreas(client, offset)
// 	if err != nil {
// 		fmt.Println("Error fetching data:", err)
// 		return
// 	}
// 	printResults(result)

// 	// Example calls:
// 	fmt.Println("Fetching next page...")
// 	nextPage(client)

// 	fmt.Println("Fetching previous page...")
// 	prevPage(client)
// }
