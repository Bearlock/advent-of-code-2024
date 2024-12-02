package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

  var left_nums []int
  var right_nums []int
  total := 0

  for scanner.Scan() {
    numbers := strings.Fields(scanner.Text())

    left_nums = append(left_nums, parse_num(numbers[0]))
    right_nums = append(right_nums, parse_num(numbers[1]))
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  sort.Ints(left_nums)
  sort.Ints(right_nums)

  for i, v := range left_nums {
    total = total + abs(v, right_nums[i])
  }

  fmt.Printf("The answer is: %d\n", total)
  return
}

func parse_num(s string) int {
  n, err := strconv.Atoi(s)

  if err != nil {
    log.Fatal(err)
  }

  return n
}

func abs (x, y int) int {
  if x > y {
    return x - y
  }

  return y - x
}
