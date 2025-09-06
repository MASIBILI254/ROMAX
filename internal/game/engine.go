package game

import (
	"fmt"
	"math/rand"
	"time"
	"ROMAX/internal/models"
	"ROMAX/pkg/utils"

)

func init() { rand.Seed(time.Now().UnixNano()) }

func pickSymbol() string {
i := utils.WeightedPick(SymbolWeights)
return Symbols[i]
}
func SpinReels() [][]string {
grid := make([][]string, Reels)
for r := 0; r < Reels; r++ {
col := make([]string, Rows)
for row := 0; row < Rows; row++ {
col[row] = pickSymbol()
}
grid[r] = col
}
return grid
}

func EvaluatePaylines(grid [][]string) (wins []models.PaylineResult, specialCount int) {
for li, line := range Paylines {
first := grid[0][line[0]]
if first == "" {
continue
}
count := 1
for r := 1; r < Reels; r++ {
s := grid[r][line[r]]
if s == first {
count++
} else {
break
}
}
if count >= 3 {
p := 0.0
if table, ok := PayoutTable[first]; ok {
if mul, ok2 := table[count]; ok2 {
p = mul
}
}
wins = append(wins, models.PaylineResult{
LineIndex: li,
Symbol: first,
Count: count,
Payout: p,
})
}
}

for r := 0; r < Reels; r++ {
for row := 0; row < Rows; row++ {
if grid[r][row] == SpecialSym {
specialCount++
}
}
}
return
}
func CascadeRefill(grid [][]string, wins []models.PaylineResult) [][]string {
for _, w := range wins {
line := Paylines[w.LineIndex]
for r := 0; r < Reels; r++ {
grid[r][line[r]] = pickSymbol()
}
}
return grid
}

func calcPayout(bet float64, pr models.PaylineResult, isFreeSpin bool) float64 {
if bet == 0 || isFreeSpin {
// if the round is a free spin, define a base bet equivalence 
base := 1.0
return base * pr.Payout
}
return bet * pr.Payout
}

func RunRound(player *models.Player, bet float64, useFreeSpin bool) (models.Round, error) {
round := models.Round{
Time: time.Now(),
Bet:  bet,
}
usingFree := false

// Handle free spin logic
if useFreeSpin {
if player.FreeSpinsRemaining <= 0 {
return round, fmt.Errorf("no free spins remaining")
}
usingFree = true
player.FreeSpinsRemaining--
bet = 0
}

// deduct bet if not a free spin
if bet > 0 {
if player.Balance < bet {
return round, fmt.Errorf("insufficient balance")
}
player.Balance -= bet
}

totalWin := 0.0
consecutive := 0
grid := SpinReels()
round.Grid = grid


for {
wins, specialCount := EvaluatePaylines(grid)
if len(wins) == 0 {
break
}


consecutive++

for _, w := range wins {
round.PaylineWins = append(round.PaylineWins, w)
p := calcPayout(bet, w, usingFree)
totalWin += p
}


// bonus trigger if specials >= 3
if specialCount >= 3 {
round.BonusTriggered = true
br := RunBonusRound()
round.BonusOutcome = &br

totalWin += br.BonusPayout
}

if award, ok := FreeSpinAwards[consecutive]; ok {
round.FreeSpinsAwarded += award
player.FreeSpinsRemaining += award
}



grid = CascadeRefill(grid, wins)
round.Grid = grid
}


round.ConsecutiveWins = consecutive
if totalWin > MaxWinCap {
totalWin = MaxWinCap
}


round.TotalWin = totalWin
player.Balance += totalWin
round.FreeSpinsRemaining = player.FreeSpinsRemaining
round.BalanceAfter = player.Balance
player.History = append(player.History, round)


return round, nil
}
