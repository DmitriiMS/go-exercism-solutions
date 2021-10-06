package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[byte]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank byte) (ret int) {
	for _, square := range cb[rank] {
		if square {
			ret++
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) (ret int) {
	for _, f := range cb {
		if file <= 8 && file > 0 && f[file-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() (ret int) {
	//ret = 64
	//return ret
	// The above will work for this case, because there's only one Chessboard
	// Lets pretend we can have custom sized Chessboards
	for _, f := range cb {
		ret += len(f)
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() (ret int) {
	for _, f := range cb {
		for _, occupied := range f {
			if occupied {
				ret++
			}
		}
	}
	return ret
}
