package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
  reqChildren := make(map[string][]string)
  total := 0

  for scanner.Scan() {
    line := scanner.Text()

    if line == "\n" {
      //do nothing
    }

    if strings.Contains(line, "|") {
      key, val := parseRule(line)
      reqChildren[key] = append(reqChildren[key], val)
    }

    if strings.Contains(line, ",") {
      pages := strings.Split(line, ",")

      if !areValidPages(pages, reqChildren) {
        sort.Slice(pages, func(i, j int) bool {
          return isParent(pages[j], reqChildren[pages[i]])
        })

        total = total + parseNum(pages[len(pages) / 2])
      }
    }
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  fmt.Printf("The answer is: %d\n", total)
}

func parseRule(rule string) (string, string) {
  keyVal := strings.Split(rule, "|")

  return keyVal[0], keyVal[1]
}

func areValidPages(pages []string, reqChildren map[string][]string) bool {
  for i, page := range pages {
    children, exists := reqChildren[page]

    if exists { // If the page isn't in the ruleset, we don't care
      for _, child := range children {
        if breaksChildRule(i, slices.Index(pages, child)) {
          return false
        }
      }
    }
  }

  return true
}

func isParent(page string, children []string) bool {
  return slices.Contains(children, page)
}

func breaksChildRule(parentIndex, childIndex int) bool {
  return childIndex != -1 && childIndex < parentIndex // -1 indicates it _isn't_ a child, but we don't care
}

func parseNum(s string) int {
  n, err := strconv.Atoi(s)

  if err != nil {
    log.Fatal(err)
  }

  return n
}

