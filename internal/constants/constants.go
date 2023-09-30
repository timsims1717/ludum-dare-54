package constants

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TestBatch = "test_batch"
	Title     = "Well that didn't work"
	Release   = 0
	Version   = 1
	Build     = 20230428

	//Initial Conditions
	InitialTrunkTargetFill = 0.5
	NumberofStrikes        = 3
)

var (
	TitleText          = "LD54"
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	TitleVariants      = []string{
		"Is this a Traveling Salesman Problem or a Packing Problem",
		"42 Days no Driving Violations",
		"Fueled by Chicken Wings and Gin",
		"First Stop after the Tetronimo Factory",
		"We need more sales! Everyone's working weekends!",
		"Your truck just barfed all over the loading dock, clean it up!",
		"We killed a bug, YAY!!!",
		"Don't forget to spend your hard earned dollars at the company store!",
		"My nose feels bigger",
		"Why do we even have that lever!?!",
		"Ain’t no fellow who regretted giving it one extra shake, but you can bet every guy has regretted giving one too few.",
	}
)

func RandomTitle() string {
	TitleText = TitleVariants[GlobalSeededRandom.Intn(len(TitleVariants))]
	return TitleText
}

type FailCondition int

const (
	RunOutOfMoney = iota
	TooManyMisses
)

func (fc FailCondition) String() string {
	switch fc {
	case RunOutOfMoney:
		return "Game Over.\nHave run out of money, you better give up and go home"
	case TooManyMisses:
		return fmt.Sprintf("Game Over.\nYou missed sales a %d vendors", NumberofStrikes)
	}
	return ""
}
