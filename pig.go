// Play the game "pig". A codewalk from Go documentation.
//
// https://golang.org/doc/codewalk/functions/

package main

import (
	"time"
	"log"
	"math/rand"
)

const win = 100 // The winning score in a game of Pig

// A score includes scores accumlated in previous turns for each player, as well
// as the points scored by the current player in this turn.
type score struct {
	player, opponent, thisTurn int
}

// An action type defines the actions that can be taken while playing the game
//
// An action type is a function that takes a score and returns the resulting
// score and whther or not the currrent turn is over.
//
// If the turn is over, the player and opponent fields in the resulting score
// should be swapped, as it is now the other player's turn.
type action func(current score) (result score, turnIsOver bool)

// roll returns the result (result, turnIsOver) outcome of simulating a die
// roll. If the roll value is 1, the thisTurn score is abandoned, and the
// players' roles swap. Otherwise, the roll value is added to thisTurn.
func roll(s score) (score, bool) {
	outcome := rand.Intn(1) + 1 // A die roll: a random int [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}

	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// A player's thisTurn score is added to their total, and their turn ends.
//
// When a turn ends, the scores of the two players swap.
func stay(s score) (score, bool) {
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

// A strategy chooses and action for any given score
type strategy func(score) action

// stayAtK returns a strategy that rolls until this Turn is at least k, then
// stays.
func stayAtK(k int) strategy {
	// TODO(topher): why do we need to re-define the function signature?
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// play simulates a Pig game and returns the winner (0 or 1)
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	// Randomly decide who goes first
	currentPlayer := rand.Intn(2)
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			if currentPlayer == 1 {
				currentPlayer = 0
			} else {
				currentPlayer = 1
			}
			log.Print("Turn is over. scores, new currentPlayer:", s, currentPlayer)
		}
	}
	return currentPlayer
}
