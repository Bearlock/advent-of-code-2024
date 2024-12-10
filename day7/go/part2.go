package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
  candidates := make([][]int, 0)
  total := 0

  for scanner.Scan() {
    parsed := make([]int, 0)
    init := strings.Split(scanner.Text(), " ")

    for i, v := range init {
      if i == 0 {
        parsed = append(parsed, parseNum(strings.Trim(v, ":")))
      } else {
        parsed = append(parsed, parseNum(v))
      }
    }

    candidates = append(candidates, parsed)
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  for _, can := range candidates {
    if canEquate(can) {
      total = total + can[0]
    }
  }

  fmt.Printf("The answer is: %d\n", total)
}

func parseNum(s string) int {
  n, err := strconv.Atoi(s)

  if err != nil {
    log.Fatal(err)
  }

  return n
}


func canEquate(equation []int) bool {
  goal := equation[0]
  nums := equation[2:]
  opLen := len(nums)
  trits := toTrits(opLen)

  for i := 0; i < trits; i++ {
    total := equation[1]
    c := strings.Split(intToBase3(i, trits, opLen), "")

    for i, v := range nums {
      if c[i] == "0" {
        total = total + v
      } else if c[i] == "1" {
        total = total * v
      } else {
        total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(v))
      }

      if total > goal {
        continue
      }
    }

    if total == goal {
      return true
    }
  }

  return false
}

func toTrits(x int) int {
  return int(math.Pow(float64(3), float64(x)))
}

func intToBase3(n, size, places int) string {
  result := ""

	if n == 0 {
		result = "0"
	} else {
    for n > 0 {
      remainder := n % 3
      result = string('0'+remainder) + result
      n /= 3
    }
  }

	return strings.Repeat("0", places - len(result)) + result
}
