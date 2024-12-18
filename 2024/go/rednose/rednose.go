package main

import(
  "fmt"
  "os"
  "bufio"
  "aoc2024.loirito/helpers"
  "strings"
  "strconv"
  "math"
)

func main(){
  file, err := os.Open(os.Args[1])
  helpers.CheckFile(file, err)
  scanner := bufio.NewScanner(file)
  defer file.Close()

  var reports [][]int
  i := 0
  for scanner.Scan(){
    elems := strings.Split(scanner.Text(), " ")
    reports = append(reports, make([]int, len(elems)))
    for j:=0; j<len(elems); j++{
      elem, _ := strconv.Atoi(elems[j])
      reports[i][j] = elem
    }
    i++
  }

  fmt.Println("Part One: ", partOne(reports))
  fmt.Println("Part Two: ", partTwo(reports))
}


func isDecreasing(report []int)(bool){
  for i:=1; i<len(report); i++{
    if report[i]-report[i-1] > 0{
      return false
    }
  }
  return true
}

func isIncreasing(report []int)(bool){
  for i:=1; i<len(report); i++{
    if report[i]-report[i-1] < 0{
      return false
    }
  }
  return true
}

func isValidReport(report []int)(bool){
  if !isDecreasing(report) && !isIncreasing(report){
    return false
  }

  for i:=1; i<len(report); i++{
    diff := report[i]-report[i-1]
    absval := math.Abs(float64(diff))
    if absval > 3 || absval < 1{
      return false
    }
  }
  
  return true
}

func partOne(input [][]int)(int){
  safe := 0
  for i:=0; i<len(input); i++{
    fmt.Println(input[i])
    if isValidReport(input[i]){
      safe++
    }
  }

  return safe
}

func partTwo(input [][]int)(int){
  safe := 0
  for i:=0; i<len(input); i++{
    if isValidReport(input[i]){
      safe++
    } else{
      for j:=0; j<len(input[i]); j++{
        if isValidReport(helpers.RemoveIntIdx(input[i], j)){
          safe++
          break
        }
      }
    }
  }

  return safe
}
