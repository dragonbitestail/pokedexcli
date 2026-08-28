# pokedex client

A command-line client to the [pokeAPI](https://www.example.com)

## Source Install

### Install dependencies

1. [Install git](https://git-scm.com/install/)
2. [Install Go Toolchain](https://go.dev/dl/)

### Retrieve

mkdir build_dir_parent

cd build_dir_parent

```
git clone https://github.com/dragonbitestail/pokedexcli.git
```

This should create pokedexcli directory in build_dir_parent.

### Build pokedexcli

```
cd pokedexcli
go build
```

You should now have an executable in `build_dir_parent/pokedexcli` called pokedexcli

### Deploy

Move pokedexcli whereever you like. Preferably in a directory in your OS's path.

### Run

If you placed pokedexcli in your OS's path:
`pokedexcli`

This will launch an interactive command prompt.

Type `help` to see available commands.
Type `exit` to return to your terminal shell.

### Environment variables

To control default logging levels where WARN is the default, use LOG_LEVEL like:
 `LOG_LEVEL=INFO pokedexcli`

where levels are those supported by the [Go slog package](https://pkg.go.dev/log/slog#Level)

To change the cache duration of retrieved and locally stored Pokedex objects use the POKEDEX_CACHE_DUR_SECS environment variable like:
`POKEDEX_CACHE_DUR_SECS=900 pokedexcli`

Combine the two:
`LOG_LEVEL=Info POKEDEX_CACHE_DUR_SECS=900 pokedexcli`
