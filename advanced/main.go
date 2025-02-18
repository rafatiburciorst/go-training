package main

import "fmt"

var dbFilmes = []string{
	"Filme 0",
	"Filme 1",
	"Filme 2",
	"Filme 3",
	"Filme 4",
	"Filme 5",
	"Filme 6",
	"Filme 7",
	"Filme 8",
	"Filme 9",
	"Filme 10",
}

func main() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// criar slice com tamanho fixo para consumir menos memória
	filmes := make([]string, 0, 10)

	// matrix2D := [][]int{}
	// matrix3D := [][][]int{}

	for _, id := range resultsFromApi {
		filme := dbFilmes[id]
		filmes = append(filmes, filme)
	}
	fmt.Println(filmes)
}
