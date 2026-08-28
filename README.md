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
