package entity

import "github.com/google/uuid"

type MatchAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewMatchAction(playerID string, minute int, action string, score int, teamID string) *MatchAction {
	return &MatchAction{
		ID:       uuid.New().String(),
		PlayerID: playerID,
		Minute:   minute,
		TeamID:   teamID,
		Action:   action,
		Score:    score}
}
