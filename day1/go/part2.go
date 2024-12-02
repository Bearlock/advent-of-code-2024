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

  // var right_nums []string
  m := make(map[string]int)
  var right_nums []string
  total := 0

  for scanner.Scan() {
    numbers := strings.Fields(scanner.Text())

    m[numbers[0]] = 0
    right_nums = append(right_nums, numbers[1])
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  for _, v := range right_nums {
    _, present := m[v]

    if present {
      m[v] = m[v] + 1
    }
  }

  for k, v := range m {
    total = total + (parse_num(k) * v)
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
