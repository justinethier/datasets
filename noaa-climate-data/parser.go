package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func ReadFile() {
  file, err := os.Open("sykesville.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan() // Skip file header
  scanner.Scan()
  for scanner.Scan() {
    ParseLine(scanner.Text())
    //fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func ParseLine(line string) {
  cols := strings.Fields(line)
  fmt.Println( len(cols) )
  fmt.Println(cols[0]) // station
  fmt.Println(cols[1]) // name
  fmt.Println(cols[9]) // date
  fmt.Println(cols[10]) // precip
}

func main() {
 ReadFile()
}
