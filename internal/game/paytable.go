package game


// Configuration for the slot: reels, rows, paylines, symbols, payouts, and free spin awards.
const (
Reels = 5
Rows = 3
MaxWinCap = 100000.0
SpecialSym = "LION"
)


var FreeSpinAwards = map[int]int{
4: 3,
5: 5,
6: 10,
7: 20,
}


// Simple payout table:
var PayoutTable = map[string]map[int]float64{
"A": {3: 5, 4: 20, 5: 50},
"K": {3: 4, 4: 15, 5: 40},
"Q": {3: 3.5, 4: 12, 5: 30},
"J": {3: 3, 4: 10, 5: 25},
"10": {3: 2.5, 4: 8, 5: 20},
"9": {3: 2, 4: 6, 5: 15},
"LION": {3: 10, 4: 50, 5: 200},
}

var Symbols = []string{"A", "K", "Q", "J", "10", "9", "LION"}
var SymbolWeights = []int{90, 100, 120, 140, 160, 180, 10}


// Paylines — 15 lines
var Paylines = [][]int{
{0, 0, 0, 0, 0},
{1, 1, 1, 1, 1},
{2, 2, 2, 2, 2},
{0, 1, 2, 1, 0},
{2, 1, 0, 1, 2},
{0, 0, 1, 0, 0},
{2, 2, 1, 2, 2},
{1, 0, 0, 0, 1},
{1, 2, 2, 2, 1},
{0, 1, 1, 1, 0},
{2, 1, 1, 1, 2},
{0, 1, 0, 1, 0},
{2, 1, 2, 1, 2},
{1, 1, 0, 1, 1},
{1, 1, 2, 1, 1},
}