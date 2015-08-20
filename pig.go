// Play the game "pig". A codewalk from Go documentation.
//
// https://golang.org/doc/codewalk/functions/

package main

import (
	"fmt"
	"math/rand"
)

const winningScore = 100 // The winning score in a game of Pig
// Number of times to play each strategy against each other one
const gamesPerSeries = 100

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
	outcome := rand.Intn(6) + 1 // A die roll: a random int [1, 6]
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
	for s.player+s.thisTurn < winningScore {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			if currentPlayer == 1 {
				currentPlayer = 0
			} else {
				currentPlayer = 1
			}
		}
	}
	return currentPlayer
}

// roundRobin simulates a series of games between every pair of strategies
//
// Returns:
//    wins: a list (of length len(strategies)) containing the number of wins for
//  each strategy
//    gamesPerStrategy: the number of games each strategy played
func roundRobin(strategies []strategy, gamesPerSeries int) (wins []int, gamesPerStrategy int) {
	wins = make([]int, len(strategies))
	for i, _ := range strategies {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				switch winner {
				case 0:
					wins[i]++
				case 1:
					wins[j]++
				}
			}
		}
	}

	// Each stragety plays gamesPerSeries times against every strategy but itself
	// (hence the '-1')
	gamesPerStrategy = gamesPerSeries * (len(strategies) - 1)

	return wins, gamesPerStrategy
}

func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	// Make one strategy for each possible staying place from 0 -> winningScore
	strategies := make([]strategy, winningScore+1)
	for i := range strategies {
		strategies[i] = stayAtK(i)
	}
	wins, gamesPerStrategy := roundRobin(strategies, gamesPerSeries)

	for i := range strategies {
		fmt.Printf("Wins, losses staying at i =% 4d: %s\n",
			i, ratioString(wins[i], gamesPerStrategy-wins[i]))
	}
}
