package models

import (
	"fmt"
	"strings"
)

type GameMode uint8

const (
	GameMode_Survival GameMode = iota
	GameMode_Normal
	GameMode_Creative
	GameMode_God
	NumGameModes
)

var (
	gameMode_Lookup = [...]string{
		"survival",
		"normal",
		"creative",
		"god",
	}
)

func (gm GameMode) String() string {
	var index int

	if gm >= NumGameModes {
		panic(gm)
	} else {
		index = int(gm)
	}

	return gameMode_Lookup[index]
}

func (gm *GameMode) EncodeWord() (string, error) {
	return gm.String(), nil
}

func (gm *GameMode) DecodeWord(str string) error {
	for i, lit := range gameMode_Lookup {
		if lit == str {
			*gm = GameMode(i)
			return nil
		}
	}

	return fmt.Errorf("realm: %s is not a valid gamemode (valid options are %s)", str, strings.Join(gameMode_Lookup[:], ", "))
}
