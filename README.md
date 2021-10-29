# Textile Tech Stack

This is a `toy` CLI application that shows the libraries and patterns used by Textile's Engineering team.

## How to build

`make build`

## Configuration

Add the `postgres-uri` connection string to the config file `~/.toy/config`. 

For testing purposes you can use the string `postgres://toy:toy@localhost/postgres?sslmode=disable&timezone=UTC` to connect to the Postgres docker container provided by `make run-postgres`.

## Usage

`./toy person`

```bash
Usage:
  toy person [command]

Available Commands:
  encode      encode receives the person info and encodes using ProtoBuffer
  list        list lists all persons info in the database
  store       store stores the person info in the database

Flags:
  -h, --help   help for person

Global Flags:
      --log-debug   Enable debug level log
      --log-json    Enable structured logging

Use "toy person [command] --help" for more information about a command.
```

### Example 1

Encode the information of a person

```bash
./toy person encode --id 1 --name Bruno
0a054272756e6f1001
```

### Example 2

Stores the information of a person in a database

```bash
./toy person store --id 2 --name Jose --email jose@gmail.com
```

### Example 3

Lits all people stored in the database

```bash
.toy person list
ID: 2, Name: Jose, Email: jose@gmail.com
```

## What is important to notice

- How `cobra` and `viper` is used
  - The use of command and sub command
  - The use of flags and persistent flags
  - The use of required flags
  - The use of the `cli` helper built by Textile
- How `ProtoBuffer` is used through `buf`
- The use of `Makefile`
- The use of `sqlc`

## Compiling `.proto` files

`make protos`

## Generating access methods for database

`make sql-assets`
