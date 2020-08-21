package main

 import (
 	"fmt"
 	"math/rand"
 	"sync"
	 "time"
 )

 type Result struct {
	 Players	string
	 Score	    int
 }

 var onlyOnce sync.Once

 var dice = []int{1, 2, 3, 4, 5, 6}

 func rollDice() int {

 	onlyOnce.Do(func() {
 		rand.Seed(time.Now().UnixNano()) 
 	})

 	return dice[rand.Intn(len(dice))]
 }

 func findMax(scores []int) (score int) {
	score = scores[0]
	for _, value := range scores {
		if value > score {
			score = value
		}
	}
	return score
}

 func main() {
	var n int
	var results []Result

	fmt.Scanln(&n)

	names := make([]string, n)
	dices := make([]int, n)
	scores := make([]int, n)

	for i:=0; i<n;i++{
		fmt.Scan(&names[i])
		dices[i] = rollDice()
	}
	
	for l:=0; l<len(dices);l++{
		for j:=0;j<=3;j++{
			
			if dices[l] == 1 || dices[l] == 3 || dices[l] == 5 {
				scores[l] = scores[l] + 10
			}else{
				scores[l] = scores[l] - 5
			}
		}
		results = append(results, Result{Players: names[l], Score: scores[l]})
	}

	fmt.Println(results)
 }