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

    if is_safe(levels, false) {
      safe_levels = safe_levels + 1
    }
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  fmt.Printf("The answer is: %d\n", safe_levels)
}

func is_safe(levels []int, ex bool) bool {
  if is_all_within_safe_bounds(levels) && is_constant_change(levels) {
    return true
  }

  if ex {
    return false
  }

  for i := range levels {
    n_levels := remove(levels, i)
    if is_safe(n_levels, true) {
      return true
    }
  }

  return false
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

func is_constant_change(l []int) bool {
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
func is_all_within_safe_bounds(l []int) bool {
  return is_all(l, is_within_safe_bounds)
}

func abs (x, y int) int {
  if x > y {
    return x - y
  }

  return y - x
}

func remove(slice []int, i int) []int {
  ret := make([]int, 0)
  ret = append(ret, slice[:i]...)
  return append(ret, slice[i+1:]...)
}
