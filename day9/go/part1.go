package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
  file, err := os.Open("../input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  disk := make([]int, 0)

  for scanner.Scan() {
    disk = mappy(strings.Split(scanner.Text(), ""), parseNum)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  files := everyOther(disk)
  frees := everyOther(disk[1:])

  answer := compact(diskSlicesToString(files, frees), reduce(0, frees, sum))
  fmt.Printf("The answer is: %d\n", answer)
}

func mappy[T, V any](ts []T, fn func(T) V) []V {
  result := make([]V, len(ts))

  for i, t := range ts {
    result[i] = fn(t)
  }

  return result
}

func mapWithIndex[T, V any](ts []T, fn func(T, int) V) []V {
  result := make([]V, len(ts))

  for i, t := range ts {
    result[i] = fn(t, i)
  }

  return result
}

func reduce[T, V any](initAcc V, ts []T, fn func(T, V) V) V {
  for _, t := range ts {
    initAcc = fn(t, initAcc)
  }

  return initAcc
}

func parseNum(s string) int {
  n, err := strconv.Atoi(s)

  if err != nil {
    log.Fatal(err)
  }

  return n
}

func everyOther[T any](ts []T) []T {
  result := make([]T, 0)

  for i, t := range ts {
    if i % 2 == 0 {
      result = append(result, t)
    }
  }

  return result
}

func diskSlicesToString(files, frees []int) []string {
  res := make([]string, 0) 

  for i, f := range files {
    for j := 0; j < f; j++ {
      res = append(res, strconv.Itoa(i))
    }

    if i < len(frees) {
      for j := 0; j < frees[i]; j++ {
        res = append(res, ".")
      }
    }
  }

  return res
}

func compact(expandDisk []string, dots int) int {
  forward := make([]string, len(expandDisk))
  reverse := make([]string, len(expandDisk))
  copy(forward, expandDisk)
  copy(reverse, expandDisk)
  slices.Reverse(reverse)

  for i, rev := range reverse {
    for j, fwd := range forward {
      if fwd == "." {
        forward[j] = rev
        reverse[i] = fwd
        break
      }
    }
  }

  mult := func(x string, y int) int {
    return parseNum(x) * y
  }

  return reduce(0, mapWithIndex(forward[:len(forward) - dots], mult), sum)
}

func sum(x, y int) int {
  return x + y
}
