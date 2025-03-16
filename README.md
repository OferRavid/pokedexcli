# Pokedex CLI

A command-line interface (CLI) Pokedex application that interacts with the PokéAPI to fetch location and Pokémon data. The application supports exploring locations, catching Pokémon, inspecting caught Pokémon, and displaying an in-game Pokedex.

## Features
✅ Explore location areas and discover Pokémon in them.   
✅ Catch Pokémon with a simulated capture mechanic.   
✅ Inspect caught Pokémon to see their stats and attributes.   
✅ View a list of all caught Pokémon.   
✅ Navigate through location areas with pagination.   

---

## Project Structure
The project is organized into the following packages:

### `main` (Root Package)
Handles the REPL (Read-Eval-Print Loop), command processing, and application logic.

- [`main.go`](https://github.com/OferRavid/pokedexcli/blob/main/main.go): Initializes the application and starts the REPL.
- [`commands.go`](https://github.com/OferRavid/pokedexcli/blob/main/commands.go): Implements the CLI commands.
- [`repl.go`](https://github.com/OferRavid/pokedexcli/blob/main/repl.go): Manages the interactive REPL interface.

### `internal/pokeapi`
Interacts with the PokéAPI to fetch Pokémon and location data.

- [`client.go`](https://github.com/OferRavid/pokedexcli/blob/main/internal/pokeapi/client.go): Defines the API client for making HTTP requests and caching responses.
- [`location_types.go`](https://github.com/OferRavid/pokedexcli/blob/main/internal/pokeapi/location_types.go): Defines data structures for locations and encounters.
- [`pokemon.go`](https://github.com/OferRavid/pokedexcli/blob/main/internal/pokeapi/pokemon.go): Defines data structures for Pokémon details.
- [`pokeapi.go`](https://github.com/OferRavid/pokedexcli/blob/main/internal/pokeapi/pokeapi.go): Contains base API URLs and constants.

### `internal/pokecache`
Implements an in-memory cache to reduce redundant API calls and improve performance.

- [`pokecache.go`](https://github.com/OferRavid/pokedexcli/blob/main/internal/pokecache/pokecache.go): Defines the Cache struct and its constructor and methods.

---

## Installation
Ensure you have Go installed, then clone the repository and build the application:

```sh
# Clone the repository
git clone https://github.com/OferRavid/pokedexcli.git
cd pokedexcli

# Build the project
go build -o pokedexcli
```

## Usage
Run the application:

```sh
./pokedexcli
```

### Available commands:

| Command   | | args...     | Description 
|-----------|-|-------------|-------------
| `help`    | |  -          | Displays a list of available commands.
| `exit`    | |  -          | Closes the application.
| `map`     | |  -          | Lists the next batch of location areas.
| `mapb`    | |  -          | Lists the previous batch of location areas.
| `explore` | |  `location` | Displays Pokémon found in the specified location.
| `catch`   | |  `pokemon`  | Attempts to catch a Pokémon from the last explored area.
| `inspect` | |  `pokemon`  | Displays details about a caught Pokémon.
| `pokedex` | |  -          | Lists all caught Pokémon.

---

## Example Session

```
$ ./pokedex
Pokedex >
Pokedex > help

Welcome to the Pokedex!
Usage:

explore <location_name>: Displays all Pokemon in the area given
catch <pokemon_name>: Attempts to catch a Pokemon encountered in an area
inspect <pokemon_name>: Shows details about a caught Pokemon
pokedex: Displays all the Pokemon you caught
exit: Exit the Pokedex
help: Displays a help message
map: Displays the next batch of 20 location-areas
mapb: Displays the previous batch of 20 location-areas

Pokedex >
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex >
Pokedex > explore sunyshore-city-area
Exploring sunyshore-city-area...
Found Pokemon: 
 - tentacool
 - tentacruel
 - staryu
...

Pokedex >
Pokedex > catch staryu
Throwing a Pokeball at staryu...
staryu was caught!
You may now inspect it with the inspect command.

Pokedex >
Pokedex > inspect staryu
Name: staryu
Height: 8
Weight: 345
Stats:
...

Pokedex >
Pokedex > exit
Closing the Pokedex... Goodbye!
```


## Dependencies
- Go standard library (for HTTP requests and CLI handling)
- [PokéAPI](https://pokeapi.co/) (for Pokémon data retrieval)

## Contributing

Feel free to fork this repository and submit pull requests! Suggestions and improvements are always welcome. 

## License
This project is open-source and available under the MIT License. See [`LICENSE`](https://github.com/OferRavid/pokedexcli/blob/main/LICENSE) for details.

---

🎮 Happy Pokémon hunting!
