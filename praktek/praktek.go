package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name   string
	dice   []int
	points int
}

type Game struct {
	players []Player
	n       int
	d       int
}

func main() {
	fmt.Print("Masukkan jumlah pemain: ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Masukkan jumlah dadu: ")
	var d int
	_, err = fmt.Scanf("%d", &d)
	if err != nil {
		fmt.Println(err)
		return
	}

	game := Game{
		players: make([]Player, n),
		n:       n,
		d:       d,
	}

	for i := 0; i < n; i++ {
		game.players[i] = Player{
			name:   fmt.Sprintf("Pemain %d", i+1),
			dice:   make([]int, d),
			points: 0,
		}
	}

	for {
		for i := 0; i < n; i++ {
			for j := 0; j < d; j++ {
				game.players[i].dice[j] = rand.Intn(6) + 1
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < d; j++ {
				if game.players[i].dice[j] == 6 {
					game.players[i].points += 6
					game.players[i].dice = append(game.players[i].dice[:j], game.players[i].dice[j+1:]...)
					j--
				}
			}

			for j := 0; j < d; j++ {
				if game.players[i].dice[j] == 1 {
					game.players[i].dice[j] = game.players[(i+1)%n].dice[0]
					game.players[(i+1)%n].dice = append(game.players[(i+1)%n].dice[1:], game.players[(i+1)%n].dice[0])
				}
			}
		}

		isFinished := true
		for i := 0; i < n; i++ {
			if len(game.players[i].dice) > 0 {
				isFinished = false
				break
			}
		}

		if isFinished {
			fmt.Println("Pemenang:", game.players[0].name)
			break
		}
	}
}
