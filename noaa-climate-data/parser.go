package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
  "strconv"
  "github.com/justinethier/keyva/lsm"
  "encoding/json"
)

var tbl = lsm.New("db", 5000)

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
  key := cols[0] + "/" + cols[1] + "/" + cols[9] // Station / Name / Date
  snowlvl, _ := strconv.Atoi(cols[11])
  if snowlvl < 0 {
    snowlvl = 0
  }
  fmt.Println(key + " - " + strconv.Itoa(snowlvl))
  b, _ := json.Marshal(snowlvl)
  tbl.Set(key, b)
}

func main() {
 ReadFile()
}
