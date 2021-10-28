# Textile Tech Stack

This is a `toy` CLI application that shows the libraries and patterns used by Textile's Engineering team.

The application reads a person's data through flags and encodes it using `ProtoBuffer` and outputs it to `stdout`.

## How to build

`make build`

## How to run

### Example 1

`./toy person encode --id 1 --name Bruno`

outputs `0a054272756e6f1001`

### Example 2

`./toy person encode --id 1 --name Bruno --email brunoangelicalza@gmail.com`

outputs `0a054272756e6f10011a1a6272756e6f616e67656c6963616c7a6140676d61696c2e636f6d`

## What is important to notice

- How `cobra` and `viper` is used
  - The use of command and sub command
  - The use of flags and persistent flags
  - The use of required flags
  - The use of the `cli` helper built by Textile
- How `ProtoBuffer` is used through `buf`
- The use of `Makefile`

## Compiling `.proto` files

`make protos`
