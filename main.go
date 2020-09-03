package main

import (
  "encoding/csv"
  "flag"
  "fmt"
  "os"
  "strings"
  "time"
)

func main () {
  csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
  timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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
  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
  correct := 0

loop:
  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = \n", i + 1, p.question)

    answerChannel := make(chan string)
    go func() {
      var answer string
      fmt.Scanf("%s", &answer)
      answerChannel <- answer
    }()

    select {
    case <-timer.C:
      fmt.Println()
      break loop
    case answer := <-answerChannel: 
      if answer == p.answer {
        correct++
      }
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
