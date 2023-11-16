package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// chessboard has key of string (ranks), and value of File
	cnt := 0
	for _, value := range cb[file] {
		if value {
			cnt++
		}
	}

	return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	invalidRank := rank < 1 || rank > 8
	if !invalidRank {
		// now iterate through all files at the specified rank and get the cnt
		for _, value := range cb {
			if value[rank-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	cols := 0
	for _, value := range cb {
		cols += len(value)
	}
	return cols
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cnt := 0
	for _, value := range cb {
		for _, val := range value {
			if val {
				cnt++
			}
		}
	}
	return cnt
}
