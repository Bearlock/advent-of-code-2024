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
  opLen := len(nums) // There are 2^n possible combinations
  bits := toBits(opLen)
  combos := make([][]bool, 0)
 
  for i := 0; i < bits; i++ {
    combos = append(combos, bitsToBoolSlice(i, toBits(opLen - 1)))
  }

  for _, c := range combos {
    total := equation[1]
    for i, v := range nums {
      if c[i] {
        total = total + v
      } else {
        total = total * v
      }
    }

    if total == goal {
      return true
    }
  }

  return false
}

func toBits(x int) int {
  return int(math.Pow(float64(2), float64(x)))
}

func bitsToBoolSlice(bits, n int) []bool {
  bs := make([]bool, n)

  for i := 0; i < n; i++ {
    bs[i] = (bits & (1 << i)) != 0
  }

  return bs
}
