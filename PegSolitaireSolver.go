package main

import "fmt"

func main() {
    p := make([]bool, 15)
	for i := 1; i < len(p); i++ {
		p[i] = true
	}

	zeroToFourteen := []*bool{&p[0], &p[2], &p[5], &p[9], &p[14]}
	oneToThirteen := []*bool{&p[1], &p[4], &p[8], &p[13]}
	threeToTwelve := []*bool{&p[3], &p[7], &p[12]}
	
	fourteenToTen := []*bool{&p[14], &p[13], &p[12], &p[11], &p[10]}
	nineToSix := []*bool{&p[9], &p[8], &p[7], &p[6]}
	fiveToThree := []*bool{&p[5], &p[4], &p[3]}
	
	tenToZero := []*bool{&p[10], &p[6], &p[3], &p[1], &p[0]}
	elevenToTwo := []*bool{&p[11], &p[7], &p[4], &p[2]}
	twelveToFive := []*bool{&p[12], &p[8], &p[5]}
	
	for {
		if tryMovePeg(zeroToFourteen) || tryMovePeg(oneToThirteen) || tryMovePeg(threeToTwelve) || 
		   tryMovePeg(fourteenToTen) || tryMovePeg(nineToSix) || tryMovePeg(fiveToThree) ||
		   tryMovePeg(tenToZero) || tryMovePeg(elevenToTwo) || tryMovePeg(twelveToFive) {
			printBoard(p)
			fmt.Println("==================")
		} else {
			break
		}
	}
	
	fmt.Println("Done!")
}

func tryMovePeg(line []*bool) (bool) {
	for i := 0; i <= len(line)-3; i++ {
		
		if !*line[i+1] || *line[i] == *line[i+2] {
			continue
		} else {
			*line[i] = !*line[i]
			*line[i+1] = !*line[i+1]
			*line[i+2] = !*line[i+2]
			
			return true
		}
	}
	return false
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