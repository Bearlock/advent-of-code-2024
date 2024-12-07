package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
  North = "n"
  South = "s"
  East = "e"
  West = "w"
  Obstacle = "#"
)

type guard struct {
  orientation string
  steps int
  done bool
  history []string
  x int
  y int
}

func (g *guard) step(gMap [][]string) {
  fmt.Println(g)
  xMax := len(gMap[0]) - 1
  yMax := len(gMap) - 1

  switch g.orientation {
  case North:
    pY := g.y - 1
    if !outOfBounds(g.x, pY, xMax, yMax) {
      if gMap[pY][g.x] == Obstacle {
        g.turn()
      } else {
        g.steps++
        g.y--
      }
    } else {
      g.done = true
    }
  case South:
    pY := g.y + 1
    if !outOfBounds(g.x, pY, xMax, yMax) {
      if gMap[pY][g.x] == Obstacle {
        g.turn()
      } else {
        g.steps++
        g.y++
      }
    } else {
      g.done = true
    }
  case East:
    pX := g.x + 1
    if !outOfBounds(pX, g.y, xMax, yMax) {
      if gMap[g.y][pX] == Obstacle {
        g.turn()
      } else {
        g.steps++
        g.x++
      }
    } else {
      g.done = true
    }
  case West:
    pX := g.x - 1
    if !outOfBounds(pX, g.y, xMax, yMax) {
      if gMap[g.y][pX] == Obstacle {
        g.turn()
      } else {
        g.steps++
        g.x--
      }
    } else {
      g.done = true
    }
  }
  // case South:
  //   if !outOfBounds(g.x, g.y + 1)
  //   if g.y + 1 <= len(gMap) - 1 && gMap[g.y + 1][g.x] == Obstacle {
  //     g.turn()
  //   } else if !outOfBounds(g.x, g.y, len(gMap[0]) - 1, len(gMap) - 1) {
  //     g.steps++
  //     g.y++
  //   }
  // case East:
  //   if g.x + 1 <= len(gMap[0]) - 1 && gMap[g.y][g.x + 1] == Obstacle {
  //     g.turn()
  //   } else if !outOfBounds(g.x, g.y, len(gMap[0]) - 1, len(gMap) - 1) {
  //     g.steps++
  //     g.x++
  //   }
  // case West:
  //   if g.x - 1 >= 0 && gMap[g.y][g.x - 1] == Obstacle {
  //     g.turn()
  //   } else if !outOfBounds(g.x, g.y, len(gMap[0]) - 1, len(gMap) - 1) {
  //     g.steps++
  //     g.x--
  //   }
  // }

  coordStr := fmt.Sprint("(", g.x,", ", g.y, ")")

  if !slices.Contains(g.history, coordStr) {
    g.history = append(g.history, coordStr)
  }
}

func (g *guard) turn() {
  turnProtocol := map[string]string {
    North: East,
    East: South,
    South: West,
    West: North,
  }

  g.orientation = turnProtocol[g.orientation]
}

func (g guard) new(avatar string, x, y int) guard {
  toOrientation := map[string]string {
    "^": North,
    ">": East,
    "v": South,
    "<": West,
  }
  coordStr := fmt.Sprint("(", x,", ", y, ")")


  return guard{
    orientation: toOrientation[avatar],

    steps: 0,
    done: false,
    x: x,
    y: y,
    history: []string{coordStr},
  }
}

func main() {
  file, err := os.Open("../input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  guardMap := make([][]string, 0)
  yIn := 0
  var thisGuard guard

  // Single flat list or list of lists? 

  for scanner.Scan() {
    corridor := strings.Split(scanner.Text(), "")
    guardMap = append(guardMap, corridor)

    for i, v := range corridor {
      if isGuard(v) {
        thisGuard = thisGuard.new(v, i, yIn)
      }
    }

    yIn++
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  for thisGuard.done != true {
    thisGuard.step(guardMap)
  }
  fmt.Println(thisGuard)

  fmt.Printf("The answer is: %d\n", len(thisGuard.history))
}

func isGuard(g string) bool {
  gs := []string{"^", ">", "v", "<"}
  return slices.Contains(gs, g)
}

func outOfBounds(x, y, xMax, yMax int) bool {
  return x > xMax || y > yMax || x < 0 || y < 0
}
