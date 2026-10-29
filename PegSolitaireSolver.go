package main

import (
    "fmt"
    "math/big"
)

var resultsByBoardState map[uint64][]big.Int
var lineList [][]int
var size int

func main() {
	size = 33
    boardStart := setupBoard(size)
			
	finalPegCount := recursiveSolve(boardStart)
	
	for pegs, count := range finalPegCount{
		if count.Sign() > 0 {
			fmt.Printf("%d pegs: %s\n", pegs, count.String())
		}
	}
}

func recursiveSolve(board uint64) []big.Int {
	//Skip solving if board position has already been seen
	if r, solved := resultsByBoardState[board]; solved {
		return r
	}
	
	result := make([]big.Int, size)
	moveMade := false
	
	//Check each line across game board
	for _, line := range lineList{
		//Check 3 spots along the line at time
		for i := 0; i+2 < len(line); i++ {
			a, b, c := uint64(1) << line[i], uint64(1) << line[i+1], uint64(1) << line[i+2]
			
			//If middle has peg and exactly one other has peg, make move
			if board&b != 0 && (board&a != 0) != (board&c != 0) {
				moveMade = true
				subResult := recursiveSolve(board ^ a ^ b ^ c)
				for endPegs := range subResult {
					result[endPegs].Add(&result[endPegs], &subResult[endPegs])
				}
			}
		}
	}
	
	//if no moves were made with this board, the game is over
	if !moveMade {
		n := 0
		//Loop counts number of occupied pegs by removing rightmost set bit each iter
		for x := board; x != 0; x &= x - 1 {
			n++
		}
		result[n].SetInt64(1)
	}
	
	resultsByBoardState[board] = result
	return result
}

func setupBoard(size int) uint64{
	emptySlot := 0
	switch size{
		case 6:
			lineList = [][]int{{0, 2, 5}, {5, 4, 3}, {3, 1, 0}}
		case 10:
			lineList = [][]int{
				{0, 2, 5, 9}, {1, 4, 8},
				{9, 8, 7, 6}, {5, 4, 3},
				{6, 3, 1, 0}, {7, 4, 2},
			}
		case 15:
			lineList = [][]int{
				{0, 2, 5, 9, 14}, {1, 4, 8, 13}, {3, 7, 12},
				{14, 13, 12, 11, 10}, {9, 8, 7, 6}, {5, 4, 3},
				{10, 6, 3, 1, 0}, {11, 7, 4, 2}, {12, 8, 5},
			}
		case 33:
			lineList = [][]int{
				{0, 1, 2}, {3, 4, 5}, {6, 7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18, 19}, {20, 21, 22, 23, 24, 25, 26},
				{27, 28, 29}, {30, 31, 32},
				{6, 13, 20}, {7, 14, 21}, {0, 3, 8, 15, 22, 27, 30},
				{1, 4, 9, 16, 23, 28, 31}, {2, 5, 10, 17, 24, 29, 32},
				{11, 18, 25}, {12, 19, 26},
			}
			emptySlot = 16
	}
	
	resultsByBoardState = map[uint64][]big.Int{}
	return ((uint64(1) << size) - 1) &^ (uint64(1) << emptySlot)
}

////////////////////////////
//                        //
//           0            //
//         1   2          //
//       3   4   5        //
//     6   7   8   9      //
//  10  11  12  13  14    //
//                        //
////////////////////////////
//						  //
//	   	  00 01 02		  //
//		  03 04 05		  //
//	06 07 08 09 10 11 12  //
//	13 14 15 16 17 18 19  //
//	20 21 22 23 24 25 26  //
//		  27 28 29		  //
//		  30 31 32		  //
//						  //
////////////////////////////