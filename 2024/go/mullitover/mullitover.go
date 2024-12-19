package main

import(
  "regexp"
  "fmt"
  "os"
  "aoc2024.loirito/helpers"
  "bufio"
  "strings"
  "strconv"
)


func main(){
  file, err := os.Open(os.Args[1])
  helpers.CheckFile(file, err)
  scanner := bufio.NewScanner(file)
  defer file.Close()
  str := ""
  for scanner.Scan(){
    str += scanner.Text()
  }
  fmt.Println(str)

  fmt.Println("rgx:")
  //partOne(str)
  partTwo(str)
}


func partOne(str string) ([]string){
  r := regexp.MustCompile(`mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)`)
  matches := r.FindAllString(str, -1)
  r = regexp.MustCompile(`[0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?`)
  res := 0
  for _, v := range matches{
    values := r.FindAllString(v, -1)
    res += multValues(values[0])
    fmt.Println(" = " , multValues(values[0]))
    fmt.Println("part one --> ", res)
    //fmt.Println(values, " = ", multValues(values))
    
  }


  return matches
}

func partTwo(str string) ([]string){
  r := regexp.MustCompile(`mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)|do(n\'t)?\(\)`)
  matches := r.FindAllString(str, -1)
  r = regexp.MustCompile(`[0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?|do(n\'t)?\(\)`)
  do := true 
  res := 0
  for _, v := range matches{
    values := r.FindAllString(v, -1)
    fmt.Println(values)
    if values[0] == "don't()"{
      do = false
    } else if values[0] == "do()"{
      do = true
    } else if do{
      res += multValues(values[0])
    }
  }

  fmt.Println("part two --> ", res)
  return matches
}

func multValues(str string)(int){
  vals := strings.Split(str,",")
  val1, _ := strconv.Atoi(vals[0])
  val2, _ := strconv.Atoi(vals[1])
  return val1 * val2
}
