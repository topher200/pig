# pig

My implementation of the Pig codewalk from the go documentation.

I'm not copy/pasting the code, but writing it by hand with my own tweaks to
learn. This initial commit includes tests, which the original did not.

## How it works

It plays the game of pig (https://en.wikipedia.org/wiki/Pig_(dice_game)) with
two players. It uses pre-defined strategies for both players. It plays many
matches of pig, pitting each of the stratagies against each other.

It outputs the percentage of wins compared to total games played for each
strategy.

## Running

Built with go version 1.4.2

Run with `go run pig.go`

## Testing

`go test`

#### Requires

`github.com/stretchr/testify/assert`
