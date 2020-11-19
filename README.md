# open-calculator
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/laupski/open-calculator?style=flat-square)
![GitHub](https://img.shields.io/github/license/laupski/open-calculator?style=flat-square)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/laupski/open-calculator?style=flat-square)
![GitHub all releases](https://img.shields.io/github/downloads/laupski/open-calculator/total?style=flat-square)

This project is an example of base `golang` usage.

![gopher](./gopher.png)

# Features
- CLI
- API

# Future Features
- gRPC
- UI
- Dockerize
- Redis Cache

# Requirements
- `go` 1.15

# Install
`git clone https://github.com/laupski/open-calculator.git` \
`cd open-calculator` \
`go build` 

To evaluate expressions such as 2+2: `./open-calculator evaluate "2+2"` \
To start the api: `./open-calculator api`

# Explanation
The fundamental logic behind this calculator is the Shunting-yard algorithm proposed by Edsger Dijkstra. Using postfix notation, the tokens become evaluated normally. More information can be found below:
- [Shunting-yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
- [Postfix Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation)