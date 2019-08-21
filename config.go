package main

var (
	// InitialDifficulty represents the starting diff
	InitialDifficulty uint = 4
	// BaselinePrecisionDigits represents the starting precision level
	BaselinePrecisionDigits uint = 128
	// MaximumPrecisionDigits represents the max precision level
	MaximumPrecisionDigits uint = 1000
	// CurrentPrecision represents the current precision level
	CurrentPrecision uint = 128
	// DiffAdjustInBlocks represents the blocks that need to pass to recalculate diff
	DiffAdjustInBlocks uint = 100
	// MaxDiffChange represents the max diff change per block
	MaxDiffChange uint = 1
	// BlockDurationTarget represents the block time target in seconds
	BlockDurationTarget uint = 12
)
