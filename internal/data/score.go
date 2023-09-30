package data

type score struct {
	Cash                 int
	SuccessfulDeliveries int
	MissedDeliveries     int
}

func NewScore() {
	CurrentScore = &score{}
}

var CurrentScore *score
