package chessboard

// type named File which stores if a square is occupied by a piece
type File []bool
// type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0
    
    if file >= "A" && file <= "H" {
        for _, val := range cb[file] {
             if val {
                 occupiedSquares++
             }
        }
    }

    return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedRank := 0
	    
    if rank >= 1 && rank <= 8 {
        for key, _ := range cb {
        	if cb[key][rank - 1] { // so ranks in Chessboard are counted from 0 to 7
            	occupiedRank++
        	}
        }
    }
    return occupiedRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
    
    for key, _ := range cb {
        totalSquares += len(cb[key])
    }

    return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalOccupiedSquares := 0

    for key, _ := range cb {
        totalOccupiedSquares += CountInFile(cb, key) // also can be replaced with CountInRank function
    }

    return totalOccupiedSquares
}
