package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type indexAndBoundary struct {
  direction string
  indexFn func(c []int, ln int) []int
  boundaryFn func(c []int) bool
}

func main() {
  file, err := os.Open("../input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  board := make([][]string, 0)
  searches := make([]string, 0)
  total := 0

  for scanner.Scan() {
    board = append(board, strings.Split(scanner.Text(), ""))
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  coords := getCoords(board, "X")

  for _, c := range coords {
    searches = append(searches, searchAll(c, board)...)
  }

  for _, s := range searches {
    if s == "XMAS" {
      total++
    }
  }

  fmt.Printf("The answer is: %d\n", total)
}

func searchAll(coords []int, board [][]string) []string {
  searches := make([]string, 0)
  xMax := len(board[0]) - 1
  yMax := len(board) - 1

  indicesAndBoundaries := map[string]indexAndBoundary{
    "north": indexAndBoundary{
      direction: "north",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[1] - index)
        coords[1] = coords[1] - index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[1] - 3)
        return coords[1] - 3 >= 0
      },
    },
    "south": indexAndBoundary{
      direction: "south",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[1] + index)
        coords[1] = coords[1] + index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[1] + 3)
        return coords[1] + 3 <=yMax
      },
    },
    "east": indexAndBoundary{
      direction: "east",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] + index)
        coords[0] = coords[0] + index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] + 3, xMax)
        return coords[0] + 3 <= xMax
      },
    },
    "west": indexAndBoundary{
      direction: "west",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] - index)
        coords[0] = coords[0] - index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] - 3)
        return coords[0] - 3 >= 0
      },
    },
    "northeast": indexAndBoundary{
      direction: "northeast",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] - index)
        coords[1] = coords[1] - index
        coords[0] = coords[0] + index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] - 3)
        return coords[1] - 3 >= 0 && coords[0] + 3 <= xMax
      },
    },
    "northwest": indexAndBoundary{
      direction: "northwest",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] - index)
        coords[1] = coords[1] - index
        coords[0] = coords[0] - index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] - 3)
        return coords[1] - 3 >= 0 && coords[0] - 3 >= 0
      },
    },
    "southeast": indexAndBoundary{
      direction: "southeast",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] - index)
        coords[1] = coords[1] + index
        coords[0] = coords[0] + index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] - 3)
        return coords[1] + 3 <= yMax && coords[0] + 3 <= xMax
      },
    },
    "southwest": indexAndBoundary{
      direction: "southeast",
      indexFn: func(coords []int, index int) []int {
        fmt.Println(coords, coords[0] - index)
        coords[1] = coords[1] + index
        coords[0] = coords[0] - index
        return coords
      },
      boundaryFn: func(coords []int) bool {
        fmt.Println(coords, coords[0] - 3)
        return coords[1] + 3 <= yMax && coords[0] - 3 >= 0
      },
    },
  }

  for _, value := range indicesAndBoundaries {
    if value.boundaryFn(coords) {
      searches = append(searches, search(coords, board, value.indexFn))
    }
  }

  return searches
}

func search(coords []int, board [][]string, indexFn func(coords []int, index int) []int) string  {
  var sb strings.Builder

  index := 0
  for index < 4 {
    copyS := make([]int, 0)
    copyS = append(copyS, coords...)
    xy := indexFn(copyS, index)
    sb.WriteString(board[xy[1]][xy[0]])
    index++
  }

  return sb.String()
}

func getCoords(board [][]string, start string) [][]int {
  coords := make([][]int, 0)

  for y := range board {
    for x, v := range board[y] {
      if v == start {
        coord := []int{x, y}
        coords = append(coords, coord)
      }
    }
  }

  return coords
}
