package game


import (
"math/rand"
"time"
"ROMAX/internal/models"
)


func init() { rand.Seed(time.Now().UnixNano()) }


// RunBonusRound simulates the Dueling Lions mini-game and returns a BonusResult
func RunBonusRound() models.BonusResult {
// Simple design: random multiplier between 1.5 and 5.0, and 20% chance of extra flat payout
mult := 1.5 + rand.Float64()*3.5
extra := 0.0
if rand.Intn(100) < 20 {
extra = 50 + rand.Float64()*150
}
return models.BonusResult{
WinMultiplier: mult,
BonusPayout: extra,
Details: "Dueling Lions outcome",
}
}