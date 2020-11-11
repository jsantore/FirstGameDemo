package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"testing"
)

func TestGot_gold(t *testing.T) {
	picture, _, err := ebitenutil.NewImageFromFile("galleon.png")
	if err != nil {
		log.Fatal("failed to load image", err)
	}
	coins, _, err := ebitenutil.NewImageFromFile("gold-coins-large.png")
	if err != nil {
		log.Fatal("failed to load image", err)
	}
	tests := []struct {
		playerStruct Sprite
		otherStruct  Sprite
		expected     bool
	}{
		{Sprite{
			pict: picture,
			xLoc: 100,
			yLoc: 100,
		},
			Sprite{pict: coins,
				xLoc: 120,
				yLoc: 100},
			true},
		{Sprite{
			pict: picture,
			xLoc: 100,
			yLoc: 100,
		},
			Sprite{pict: coins,
				xLoc: 500,
				yLoc: 500},
			false},
		{Sprite{
			pict: coins,
			xLoc: 0,
			yLoc: 0,
		},
			Sprite{pict: coins,
				xLoc: 0,
				yLoc: 0},
			true},
		{Sprite{
			pict: coins,
			xLoc: -math.MaxInt16,
			yLoc: -math.MaxInt16,
		},
			Sprite{pict: coins,
				xLoc: -math.MaxInt16,
				yLoc: -math.MaxInt16 + 5},
			true},
	}
	for _, testCase := range tests {
		result := gotGold(testCase.playerStruct, testCase.otherStruct)
		if result != testCase.expected {
			t.Error("collision Detection Error in test case", testCase)
		}
	}
}
