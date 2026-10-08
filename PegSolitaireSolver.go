package main

import "fmt"

var p []bool
var lineList [][]*bool
var finalPegCount []int

func main() {
    setupBoard(15)
	
	printBoard(p)
		
	recursiveSolve(p)
	
	fmt.Println(finalPegCount)
}

func recursiveSolve(p []bool) {
	moveMade := false
	for _, line := range lineList{
		//Search for moves
		moves := findMoves(line)
		
		//For each possible move, make move, then call self with new board, then reverse move
		for _, i := range moves{
			makeMove(line, i)
			recursiveSolve(p)
			makeMove(line, i)
			
			moveMade = true
		}
	}
	
	//if no moves were made with this board, send board state to channel
	if !moveMade {
		foundEndState(p)
	}
	return
}

func findMoves(line []*bool) ([]int) {
	moves := []int{}
	for i := 0; i <= len(line)-3; i++ {
		if *line[i+1] && *line[i] != *line[i+2] {
			moves = append(moves, i)
		} 
	}
	return moves
}

func makeMove(line []*bool, i int) {
	*line[i] = !*line[i]
	*line[i+1] = !*line[i+1]
	*line[i+2] = !*line[i+2]
}

func foundEndState(pegs []bool){
	pegCount := 0
	for _, v := range pegs{
		if v {
			pegCount++
		}
	}
	
	finalPegCount[pegCount]++
}

func printBoard(pegs []bool){
	for i, v := range pegs{
		char := "O"
		if v {
			char = "X"
		}
		
		switch i{
			case 0, 2, 5, 9, 14:
				fmt.Println(char)
			default:
				fmt.Printf("%s ", char)
		}
	}
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
func setupBoard(size int){
	p = make([]bool, size)
	finalPegCount = make([]int, size)
	
	switch size{
		case 6:
			var zeroToFive = []*bool{&p[0], &p[2], &p[5]}
			var fiveToThree = []*bool{&p[5], &p[4], &p[3]}
			var threeToZero = []*bool{&p[3], &p[1], &p[0]}

			lineList = [][]*bool{zeroToFive, fiveToThree, threeToZero}
		case 10:
			var zeroToNine = []*bool{&p[0], &p[2], &p[5], &p[9]}
			var oneToEight = []*bool{&p[1], &p[4], &p[8]}

			var nineToSix = []*bool{&p[9], &p[8], &p[7], &p[6]}
			var fiveToThree = []*bool{&p[5], &p[4], &p[3]}

			var sixToZero = []*bool{&p[6], &p[3], &p[1], &p[0]}
			var sevenToTwo = []*bool{&p[7], &p[4], &p[2]}

			lineList = [][]*bool{zeroToNine, oneToEight, nineToSix, fiveToThree, sixToZero, sevenToTwo}
		case 15:
			var zeroToFourteen = []*bool{&p[0], &p[2], &p[5], &p[9], &p[14]}
			var oneToThirteen = []*bool{&p[1], &p[4], &p[8], &p[13]}
			var threeToTwelve = []*bool{&p[3], &p[7], &p[12]}

			var fourteenToTen = []*bool{&p[14], &p[13], &p[12], &p[11], &p[10]}
			var nineToSix = []*bool{&p[9], &p[8], &p[7], &p[6]}
			var fiveToThree = []*bool{&p[5], &p[4], &p[3]}

			var tenToZero = []*bool{&p[10], &p[6], &p[3], &p[1], &p[0]}
			var elevenToTwo = []*bool{&p[11], &p[7], &p[4], &p[2]}
			var twelveToFive = []*bool{&p[12], &p[8], &p[5]}

			lineList = [][]*bool{zeroToFourteen, oneToThirteen, threeToTwelve, fourteenToTen, nineToSix, fiveToThree, tenToZero, elevenToTwo, twelveToFive}
	}
	
	for i := 1; i < len(p); i++ {
		p[i] = true
	}
}