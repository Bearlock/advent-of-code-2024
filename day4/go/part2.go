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

  xMax := len(board[0]) - 1
  yMax := len(board) - 1
  coords := getCoords(board, "A")

  for _, c := range coords {
    if isWithinBoundaries(c, xMax, yMax) {
      searches = append(searches, search(c, board))
    }
  }

  for _, s := range searches {
    if isValidSearch(s) {
      total++
    }
  }

  fmt.Printf("The answer is: %d\n", total)
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

func search(coords []int, board [][]string) string  {
  var sb strings.Builder

  // Top left
  x := coords[0]
  y := coords[1]

  // Top left
  sb.WriteString(board[y - 1][x - 1])

  // Top right
  sb.WriteString(board[y - 1][x + 1])

  // Bottom left
  sb.WriteString(board[y + 1][x - 1])

  //Bottom right
  sb.WriteString(board[y + 1][x + 1])
  return sb.String()
}

func isWithinBoundaries(coords []int, xMax, yMax int) bool {
  x := coords[0]
  y := coords[1]

  return x - 1 >= 0 && x + 1 <= xMax && y - 1 >= 0 && y + 1 <= yMax
}

func isValidSearch(s string) bool {
  return s == "MSMS" || s == "MMSS" || s == "SMSM" || s == "SSMM"
}
