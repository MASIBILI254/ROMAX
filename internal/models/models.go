package models

import "time"

type Player struct {
	ID                 string    `json:"id"`
	Balance            float64   `json:"balance"`
	FreeSpinsRemaining int       `json:"free_spins_remaining"`
	History            []Round   `json:"history"`
	CreatedAt          time.Time `json:"created_at"`
}

type RoundRequest struct {
	PlayerID    string  `json:"player_id"`
	Bet         float64 `json:"bet"`
	UseFreeSpin bool    `json:"use_free_spin"`
}
type Round struct {
	Time               time.Time       `json:"time"`
	Bet                float64         `json:"bet"`
	Grid               [][]string      `json:"grid"`
	PaylineWins        []PaylineResult `json:"payline_wins"`
	ConsecutiveWins    int             `json:"consecutive_wins"`
	TotalWin           float64         `json:"total_win"`
	FreeSpinsAwarded   int             `json:"free_spins_awarded"`
	FreeSpinsRemaining int             `json:"free_spins_remaining"`
	BonusTriggered     bool            `json:"bonus_triggered"`
	BonusOutcome       *BonusResult    `json:"bonus_outcome,omitempty"`
	BalanceAfter       float64         `json:"balance_after"`
}
type PaylineResult struct {
	LineIndex int     `json:"line_index"`
	Symbol    string  `json:"symbol"`
	Count     int     `json:"count"`
	Payout    float64 `json:"payout"`
}

type BonusResult struct {
	WinMultiplier float64 `json:"win_multiplier"`
	BonusPayout   float64 `json:"bonus_payout"`
	Details       string  `json:"details"`
}
