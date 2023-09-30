package constants

import (
	_ "github.com/faiface/pixel"
)

type Achievement struct {
	Name                   string
	LabelText              string
	Description            string
	Achieved               bool
	MyFamily               AchievementFamily
	AchievementFamilyOrder int
	Properties             map[string]string
	Presented              bool
}

type AchievementFamily struct {
	Name string
	//StickyNote         *object.Object
	//StickyNotePosition pixel.Vec
}

func (af *AchievementFamily) String() string {
	maxIter := -1
	message := ""
	for _, value := range Achievements {
		if value.MyFamily.Name == af.Name && value.Achieved && value.AchievementFamilyOrder > maxIter {
			maxIter = value.AchievementFamilyOrder
			message = value.LabelText
		}
	}
	return message
}
func (af *AchievementFamily) Achieved() bool {
	for _, value := range Achievements {
		if value.MyFamily.Name == af.Name && value.Achieved {
			return true
		}
	}
	return false
}

var (
	AchievementFamilies = map[string]AchievementFamily{
		//"CreateTetrominos": {
		//	Name:               "CreateTetrominos",
		//	StickyNote:         nil,
		//	StickyNotePosition: pixel.V(-40, 510),
		//},
	}

	Achievements = map[string]Achievement{
		//"Create5Tetrominos": {
		//	Name:                   "Create5Tetrominos",
		//	LabelText:              "You have met the initial quota of 5 Tetrominos, if you fall behind, it's game over!\n-Management",
		//	Description:            "Construct 5 Valid Tetrominos and deliver them to the Board",
		//	MyFamily:               AchievementFamilies["CreateTetrominos"],
		//	AchievementFamilyOrder: 0,
		//	Properties:             map[string]string{"target": "5"},
		//},

	}
)
