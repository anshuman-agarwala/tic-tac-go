package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type board struct {
	board [][]string
	turn  string
}

func NewBoard() *board {
	b := board{}
	b.board = [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	b.turn = "X"
	return &b
}

func getPossibleMoves(b board) ([][][]string, [][]int) {
	player := b.turn
	var moves [][][]string
	var moveIndex [][]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardcopy := make([][]string, 3)
			for k := 0; k < 3; k++ {
				boardcopy[k] = make([]string, 3)
				copy(boardcopy[k], b.board[k])
			}
			if boardcopy[i][j] == " " {
				moveIndex = append(moveIndex, []int{i, j})
				boardcopy[i][j] = player
				moves = append(moves, boardcopy)
			}
		}
	}
	return moves, moveIndex
}

func minimax(b board) (float64, []int) {
	winner := checkWin(b)
	if winner != "" {
		if winner == "X" {
			// fmt.Println("X wins")
			// printBoard(b.board)
			// fmt.Println("X wins")
			return 1, nil
		} else if winner == "O" {
			// fmt.Println("O wins")
			return -1, nil
		} else if winner == " " {
			return 0, nil
		}
	}

	possibleMoves, moveIndex := getPossibleMoves(b)
	bestMove := make([]int, 0)
	if b.turn == "X" {
		maxEval := math.Inf(-1)
		for index, item := range possibleMoves {
			tempBoard := board{item, "O"}
			eval, move := minimax(tempBoard)
			if eval >= maxEval {
				maxEval = eval
				bestMove = moveIndex[index]
				if move == nil {
					// printBoard(tempBoard.board)
					break
				}
			}
		}
		return maxEval, bestMove
	} else {
		minEval := math.Inf(1)
		for index, item := range possibleMoves {
			tempBoard := board{item, "X"}
			eval, move := minimax(tempBoard)
			if eval <= minEval {
				minEval = eval
				bestMove = moveIndex[index]
				if move == nil {
					// printBoard(tempBoard.board)
					break
				}
			}
		}
		return minEval, bestMove
	}

}

func checkWin(b board) string {
	for i := 0; i < 3; i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] != " " {
			return b.board[i][0]
		}
		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] != " " {
			return b.board[0][i]
		}
	}
	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[1][1] != " " {
		return b.board[0][0]
	}
	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[1][1] != " " {
		return b.board[0][2]
	}
	isDraw := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.board[i][j] == " " {
				isDraw = false
			}
		}
	}
	if isDraw {
		return " "
	}
	return ""
}

func PlayGame() {
	b := NewBoard()
	opponent := choose_opponent()
	var xInp, yInp int
	for {
		printBoard(b.board)
		if opponent == 1 || (b.turn == "X" && opponent == 2) || (b.turn == "O" && opponent == 3) {
			xInp, yInp = getInput(*b)
		} else {
			_, move := minimax(*b)
			xInp = move[0]
			yInp = move[1]
		}
		b.board[xInp][yInp] = b.turn
		result := checkWin(*b)
		if result == " " {
			printBoard(b.board)
			fmt.Println("The game has ended in a draw.")
			return
		}
		if result != "" {
			printBoard(b.board)
			fmt.Printf("%s is the winner!", b.turn)
			return
		}
		if b.turn == "X" {
			b.turn = "O"
		} else {
			b.turn = "X"
		}
	}
}

func getInput(b board) (int, int) {
	reader := bufio.NewReader(os.Stdin)
	var X, Y int
	for {
		fmt.Print("Enter X coordinate.")
		x, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while getting input. Try again.")
			continue
		}
		x = strings.Replace(x, "\n", "", -1)
		x = strings.Replace(x, "\r", "", -1)
		X, err = strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while getting input. Try again.")
			continue
		}
		fmt.Print("Enter Y coordinate.")
		y, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while getting input. Try again.")
			continue
		}
		y = strings.Replace(y, "\n", "", -1)
		y = strings.Replace(y, "\r", "", -1)
		Y, err = strconv.Atoi(y)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while getting input. Try again.")
			continue
		}
		if X > 2 || Y > 2 || X < 0 || Y < 0 || b.board[X][Y] != " " {
			fmt.Println("Invalid input entered. Try again.")
			continue
		}
		break
	}
	return X, Y
}

func printBoard(board [][]string) {
	fmt.Println("---------")
	for ix := 0; ix < 3; ix++ {
		for iy := 0; iy < 3; iy++ {
			fmt.Print("|", board[ix][iy], "|")
		}
		fmt.Println("\n---------")
	}
}

func choose_opponent() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Do you want to play against a human or the computer?")
		fmt.Println("1: Human")
		fmt.Println("2: Computer")
		choice, err := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)
		choice = strings.Replace(choice, "\r", "", -1)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while getting input. Try again.")
			continue
		}
		if choice == "1" || choice == "2" {
			intChoice, err := strconv.Atoi(choice)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error while getting input. Try again.")
				continue
			}
			if intChoice == 2 {
				fmt.Println("Who goes first?")
				fmt.Println("1: You")
				fmt.Println("2: Computer")
				firstPlayer, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error while getting input. Try again.")
					continue
				}
				firstPlayer = strings.Replace(firstPlayer, "\n", "", -1)
				firstPlayer = strings.Replace(firstPlayer, "\r", "", -1)
				if firstPlayer == "1" {
					return 2
				} else if firstPlayer == "2" {
					return 3
				} else {
					fmt.Println("Invalid Input. Try again.")
					continue
				}
			}
			return intChoice
		} else {
			fmt.Println("Invalid Input. Try again.")
		}
	}
}

func main() {
	PlayGame()
	// boarder := NewBoard()
	// x, _ := getPossibleMoves(*boarder)
	// for _, item := range x {
	// 	// fmt.Print(y[index])
	// 	fmt.Print(minimax(board{item, "X"}))
	// 	printBoard(item)
	// }

}
