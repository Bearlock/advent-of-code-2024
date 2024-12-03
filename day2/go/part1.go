package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

  safe_levels := 0

  for scanner.Scan() {
    levels := to_int_slice(strings.Fields(scanner.Text()))

    if is_safe(levels) {
      safe_levels = safe_levels + 1
    }
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  fmt.Printf("The answer is: %d\n", safe_levels)
}

func is_safe(levels []int) bool {
  return is_all(levels, is_within_safe_bounds) && is_increasing_or_decreasing(levels)
}

func to_int_slice(s_slice []string) []int {
  var levels []int

  for _, s := range s_slice {
    levels = append(levels, parse_num(s))
  }

  return levels
}

func parse_num(s string) int {
  n, err := strconv.Atoi(s)

  if err != nil {
    log.Fatal(err)
  }

  return n
}

func is_increasing(x, y int) bool {
  return x < y
}

func is_decreasing(x, y int) bool {
  return x > y
}

func is_same(x, y int) bool {
  return x == y
}

func is_all(l []int, f func(x, y int) bool) bool {
  ln := (len(l) - 1)

  for i, v := range l {
    if i < ln {
      if !f(v, l[i + 1]) {
        return false
      }
    }
  }

  return true
}

func is_increasing_or_decreasing(l []int) bool {
  x := l[0]
  y := l[1]

  if is_same(x, y) {
    return false
  }

  if is_increasing(x, y) {
    return is_all(l, is_increasing)
  }

  return is_all(l, is_decreasing)
}

func is_within_safe_bounds(x, y int) bool {
  abs_val := abs(x, y)
  return abs_val > 0 && abs_val < 4
}

func abs (x, y int) int {
  if x > y {
    return x - y
  }

  return y - x
}
