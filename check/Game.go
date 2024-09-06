package main

import (
	"fmt"
	"math/rand"
)

// iam with sador on multiway
var (
	show=fmt.Println
	playerx = "X" // we have 2 players O and X
	playero = "O"
	empty   = " " //we will dcleare this
	floor   [3][3]string
)

func main() {
	initalize_Floor()
	print_Floor()
	for !isstillGame(){
		move_Playero(playero)
	move_Playerx(playerx)
	}
	
}

// this function  initalizes every floor element as an empty at the begnning 
func initalize_Floor() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			floor[i][j] = empty //every cell is empty at the begining

		}

	}
}

// this function prints all the floor elements to be shown by the players
func print_Floor() {
	show("================")
	for i := 0; i < 3; i++ {
		print(" | ")
		for j := 0; j < 3; j++ {

			print(floor[i][j]+" | ")

		}
		show()
		show("===============")
	}

}


//this asks for playerx(Human)  row and column  and then move to that 
func move_Playerx( player string ) {
	var col,row int 
	show("plase enter the row and column you want to move player X ")
	fmt.Scan(&row,&col)
	floor[row][col]=player
	print_Floor()


}
//this asks if the player wants to play with human or computer
func move_Playero( player string ){
var col,row,choose_player int 

show("1: Play with Computer")
show("2: play with Human")
show("plase choose with whom do you want to play")
fmt.Scan(&choose_player)

if choose_player==1{ // player playing with computer
	col=rand.Intn(3)
	row=rand.Intn(3)
	floor[row][col]=player
	print_Floor()
} else if choose_player==2{
	//2 human players 
	show("plase enter the row and column you want to move playerO ")
	fmt.Scan(&row,&col)
	floor[row][col]=player
	print_Floor()

}else{
	show(" please insert only 1 and 2 ")
}


}

func isstillGame() bool{
	return findWinner(playerx)||findWinner(playero)||floorfull()
}

func floorfull() bool{
	for i := 0; i < 3; i++ {
		print(" | ")
		for j := 0; j < 3; j++ {

    if floor[i][j]==empty{
		return false 
	}
	

		}
	}

	return true

}

func findWinner(player string) bool{
	for i := 0; i < 3; i++ {
        if floor[i][0] == player && floor[i][1] == player && floor[i][2] == player {
            return true
        }
        if floor[0][i] == player && floor[1][i] == player && floor[2][i] == player {
            return true
        }
    }
   
// we have two digonals and for those 

if floor[0][0] == player && floor[1][1] == player && floor[2][2] == player {
	return true
}
if floor[0][2] == player && floor[1][1] == player && floor[2][0] == player {
	return true
}
return false
}


