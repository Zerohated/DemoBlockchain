# Simplist Blockchain Turorial
Following the turorial from [Coral Health](https://coral.health/) on Medium with slightly modification.
> Code your own blockchain mining algorithm in Go!
> [Link Here] (https://medium.com/@mycoralhealth/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f)

## Quick Start

### Get relative packages
`Iris` has some problem when using along with `dep`. All dependency need to be maintain manually.
- Web server: 
   > `go get -u github.com/kataras/iris`
- To load variable from '.env' file:
   > `go get -u github.com/joho/godotenv`
- Pretty printer for Go data structures:
   > `go get -u github.com/davecgh/go-spew/spew`
### Start localhost

1. `go run main.go`
