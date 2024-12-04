package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
  file, err := os.Open("../input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
  matches := make([]string, 0)

  for scanner.Scan() {
    matches = append(matches, pattern.FindAllString(scanner.Text(), -1)...)
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  answer := reduce(0, mappy(matches, parseAndMultiply), add)

  fmt.Printf("Answer is: %d\n", answer)
}

func parseAndMultiply(toToken string) int {
  pattern := regexp.MustCompile(`\d{1,3}`)
  pair := mappy(pattern.FindAllString(toToken, -1), parseNum)

  if len(pair) != 2 {
    log.Fatal("Somehow not getting a pair of numbers")
  }

  return pair[0] * pair[1]
}

func mappy[T, V any](ts []T, fn func(T) V) []V {
  result := make([]V, len(ts))
  for i, t := range ts {
    result[i] = fn(t)
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

func add(x, y int) int {
  return x + y
}
