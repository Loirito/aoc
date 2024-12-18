package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "aoc2024.loirito/helpers"
)


func main() {

  file, err := os.Open(os.Args[1])
  helpers.CheckFile(file, err)
  scanner := bufio.NewScanner(file)
  defer file.Close()

  var left []int
  var right []int

  for scanner.Scan(){
    elements := strings.Split(scanner.Text(), "   ")
    left_elem, _ := strconv.Atoi(elements[0])
    right_elem, _ := strconv.Atoi(elements[1])
    left = append(left, left_elem)
    right = append(right, right_elem)
  }

  left = helpers.QuicksortStart(left)
  right = helpers.QuicksortStart(right)

  res := 0

  for i:=0; i<len(left); i++{
    if left[i] < right[i]{
      res += right[i] - left[i]
    } else{
      res += left[i] - right[i]
    }
  }

  var m map[int]int
  m = make(map[int]int)

  similarity := 0

  for i:=0; i<len(left); i++{
    elem, ok := m[left[i]]
    if ok {
      similarity += left[i] * elem
    } else {
      occs := 0
      for j:=0; j<len(right) && right[j] <= left[i]; j++{
        if left[i] == right[j]{
          occs += 1
        }
      }
      m[left[i]] = occs
      similarity += left[i] * occs
    }
  }


  fmt.Println(m)
  fmt.Println(similarity)


}
