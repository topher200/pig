package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var originalScore = score{10, 50, 5}

func TestStay(t *testing.T) {
	resultingScore, turnIsOver := stay(originalScore)

	assert.True(t, turnIsOver, "Turn should be over")
	confirmTurnOver(originalScore, resultingScore, t)

	// Make sure our player's score was added up correctly
	expectedOpponentScore := originalScore.player + originalScore.thisTurn
	assert.Equal(t, expectedOpponentScore, resultingScore.opponent,
		"Resulting score of the [now] opponent should be the original total "+
			"plus their score from last turn")
}

func TestRoll(t *testing.T) {
	// Either the turn ended and we have swapped, or our thisTurn score has
	// increased.
	resultingScore, turnIsOver := roll(originalScore)

	if turnIsOver {
		confirmTurnOver(originalScore, resultingScore, t)
		assert.Equal(t, originalScore.player, resultingScore.opponent,
			"New opponent score should be our original player score")

	} else {
		assert.Equal(t, originalScore.opponent, resultingScore.opponent,
			"opponent score should not have changed. was %d, now %d",
			originalScore.opponent, resultingScore.opponent)
		// The new thisTurn should be [1, 6] more than the original
		assert.True(t, resultingScore.thisTurn > originalScore.thisTurn)
		assert.True(t, resultingScore.thisTurn <= originalScore.thisTurn+6)
	}
}

func confirmTurnOver(originalScore, resultingScore score, t *testing.T) {
	if resultingScore.thisTurn != 0 {
		t.Errorf("Turn should be over, so score for this turn should be 0")
	}

	if resultingScore.player != originalScore.opponent {
		t.Errorf("Player score should now be original opponent score. Expected %d, got %d",
			originalScore.opponent, resultingScore.player)
	}
}
