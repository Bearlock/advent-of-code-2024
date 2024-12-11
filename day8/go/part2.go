package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
  file, err := os.Open("../input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  freqToCoords := make(map[string][][]int, 0)
  y := 0

  for scanner.Scan() {
    for x, v := range strings.Split(scanner.Text(), "") {
      if v != "." {
        freqToCoords[v] = append(freqToCoords[v], []int{x, y})
      }
    }
    y++
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  antinodes := calcAntinodes(freqToCoords, y)
  fmt.Printf("The answer is: %d\n", len(antinodes))
}

func calcAntinodes(freqToCoords map[string][][]int, mVal int) [][]int {
  antinodes := make([][]int, 0)


  for key, coords := range freqToCoords {
    for i := 0; i < len(freqToCoords[key]); i++ {
      for _, coord := range coords {
        antinodes = append(antinodes, genAntiNodes(coords[i], coord, mVal)...)
      }
    }
  }
  
  f := func(s []int) bool {
    return s != nil
  }


  filtered := filter(antinodes, f)

  fmt.Println(dedupe(filtered))

  return dedupe(filtered)
}

func genAntiNodes(firstCoord, secondCoord []int, mVal int) [][]int {
  antinodes := make([][]int, 2)
  fX := firstCoord[0]
  fY := firstCoord[1]
  sX := secondCoord[0]
  sY := secondCoord[1]
  isLessX := false
  isLessY := false

  if slices.Compare(firstCoord, secondCoord) == 0 {
    return nil
  }

  if fX < sX {
    isLessX = true
  }

  if fY < sY {
    isLessY = true
  }

  dX := abs(fX - sX)
  dY := abs(fY - sY)

  antinodes = append(antinodes, []int{fX, fY}, []int{sX, sY})

  if isLessX && isLessY { // firstCoord is u-left
    fs := []int{fX - dX, fY - dY}
    ss := []int{sX + dX, sY + dY}

    for inBounds(fs, mVal) {
      antinodes = append(antinodes, fs)
      fs = []int{fs[0] - dX, fs[1] - dY}
    }

    for inBounds(ss, mVal) {
      antinodes = append(antinodes, ss)
      ss = []int{ss[0] + dX, ss[1] + dY}
    }
  } else if !isLessX && isLessY {// firstCoord is u-right
    fs := []int{fX + dX, fY - dY}
    ss := []int{sX - dX, sY + dY}

    for inBounds(fs, mVal) {
      antinodes = append(antinodes, fs)
      fs = []int{fs[0] + dX, fs[1] - dY}
    }

    for inBounds(ss, mVal) {
      antinodes = append(antinodes, ss)
      ss = []int{ss[0] - dX, ss[1] + dY}
    }
  } else if !isLessX && !isLessY {// firstCoord is d-right
    fs := []int{fX + dX, fY + dY}
    ss := []int{sX - dX, sY - dY}

    for inBounds(fs, mVal) {
      antinodes = append(antinodes, fs)
      fs = []int{fs[0] + dX, fs[1] + dY}
    }

    for inBounds(ss, mVal) {
      antinodes = append(antinodes, ss)
      ss = []int{ss[0] - dX, ss[1] - dY}
    }
  } else {
    fs := []int{fX - dX, fY + dY}
    ss := []int{sX + dX, sY - dY}

    for inBounds(fs, mVal) {
      antinodes = append(antinodes, fs)
      fs = []int{fs[0] - dX, fs[1] + dY}
    }

    for inBounds(ss, mVal) {
      antinodes = append(antinodes, ss)
      ss = []int{ss[0] + dX, ss[1] - dY}
    }
  }

  return antinodes
}

func abs(x int) int {
  if x < 0 {
    return -x
  }

  return x
}

func filter(sli [][]int, f func([]int) bool) [][]int {
  filtered := make([][]int, 0)
  for _, s := range sli {
    if f(s) {
      filtered = append(filtered, s)
    }
  }

  return filtered
}

func inBounds(s []int, mVal int) bool {
  return s[0] >= 0 && s[1] >= 0 && s[0] < mVal && s[1] < mVal
}

func dedupe(sli [][]int) [][]int {
  deduper := make(map[string][]int, 0)
  dedupes := make([][]int, 0)

  for _, s := range sli {
    deduper[fmt.Sprint(s[0], s[1])] = s
  }

  for _, v := range deduper {
    dedupes = append(dedupes, v)
  }

  return dedupes
}
