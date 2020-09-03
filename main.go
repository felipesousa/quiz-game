package main

import (
  "encoding/csv"
  "flag"
  "fmt"
  "os"
  "strings"
)

func main () {
  csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
  flag.Parse()

  file, err := os.Open(*csvFileName)
  if err != nil {
    exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFileName))
  }

  r := csv.NewReader(file)
  records, err := r.ReadAll()

  if err != nil {
    exit("Failed to read the CSV file")
  }

  problems := parseRecords(records)
  correct := 0

  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = \n", i + 1, p.question)
    var answer string
    fmt.Scanf("%s", &answer)

    if answer == p.answer {
      correct++
    }
  }

  fmt.Printf("You scored %d out of %d. \n", correct, len(records))
}

func parseRecords(records [][]string) []Record {
  data := make([]Record, len(records))
  for i, r := range records {
    data[i] = Record{
      question: r[0],
      answer: strings.TrimSpace(r[1]),
    }
  }

  return data
}

type Record struct {
  question string
  answer string
}

func exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}
