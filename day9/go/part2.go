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

type fileOrFree struct {
  typ string
  id int
  chunk []string
}

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

  answer := compact(diskSlicesToStruct(files, frees))
  fmt.Printf("The answer is: %d\n", answer)
}

func mappy[T, V any](ts []T, fn func(T) V) []V {
  result := make([]V, len(ts))

  for i, t := range ts {
    result[i] = fn(t)
  }

  return result
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

func diskSlicesToStruct(files, frees []int) []fileOrFree{
  res := make([]fileOrFree, 0)
  for i, f := range files {
    fileChunk := make([]string, 0)

    for j := 0; j < f; j++ {
      fileChunk = append(fileChunk, strconv.Itoa(i))
    }

    file := fileOrFree{
      typ: "file",
      id: i,
      chunk: fileChunk,
    }

    res = append(res, file)

    if i < len(frees) && frees[i] > 0 {
      freeChunk := make([]string, 0)

      for j := 0; j < frees[i]; j++ {
        freeChunk = append(freeChunk, ".")
      }

      free := fileOrFree{
        typ: "free",
        id: -1,
        chunk: freeChunk,
      }

      res = append(res, free)
    }
  }

  return res
}

func compact(expandStructs []fileOrFree) int {
  forward := make([]fileOrFree, len(expandStructs))
  reverse := make([]fileOrFree, len(expandStructs))
  copy(forward, expandStructs)
  copy(reverse, expandStructs)
  slices.Reverse(reverse)

  i := 0
  j := 0

  for i < len(reverse) {
    for j < len(forward) {

      fwd := forward[j]
      rev := reverse[i]

      if len(forward) - i > j && fwd.typ == "free" && len(reverse[i].chunk) <= len(forward[j].chunk) && rev.typ != "free" {
        forward[j] = rev

        if len(rev.chunk) != len(fwd.chunk) { 

          newFreeChunk := make([]string, 0)
          for nfc := 0; nfc < len(rev.chunk); nfc++{
            newFreeChunk = append(newFreeChunk, ".")
          }

          newFree := fileOrFree{
            typ: "free",
            id: -1,
            chunk: newFreeChunk,
          }

          forward[(len(forward) - 1) - i] = newFree

          appendFreeChunk := make([]string, 0)
          for afc := 0; afc < len(fwd.chunk) - len(rev.chunk); afc++{
            appendFreeChunk = append(appendFreeChunk, ".")
          }

          appendFree := fileOrFree{
            typ: "free",
            id: -1,
            chunk: appendFreeChunk,
          }

          forward = slices.Insert(
            forward,
            j + 1,
            appendFree,
          )
        } else {
          forward[len(forward) - 1 - i] = fwd
        }

        revvy := make([]fileOrFree, len(forward))
        for i := range forward {
          revvy[i] = forward[i]
        }

        slices.Reverse(revvy)
        reverse = revvy

        j++
        break 
      }
      j++
    }
    i++
    j = 0
  }

  total := 0
  noChunks := make([]int, 0)

  for _, f := range forward {
    chunkLen := len(f.chunk)
    if chunkLen > 1 {
      for j := 0; j < chunkLen; j++ {
        noChunks = append(noChunks, f.id)
      }
    } else {
      noChunks = append(noChunks, f.id)
    }
  }

  for i, c := range noChunks {
    if c != -1 {
      total = total + (i * c)
    }
  }

  return total
}

func sum(x, y int) int {
  return x + y
}
