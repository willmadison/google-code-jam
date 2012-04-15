package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
  "flag"
)

var inputFileName *string = flag.String("file", "", "The input file for this script")

type RecycledNumPair struct {
  number, recycled int
}

func main() {
  flag.Parse()

  if *inputFileName == "" {
    panic("No Input File Provided!")
  }

  inputFile, openError := os.Open(*inputFileName)

  if openError != nil {
    errorDesc := fmt.Sprint("Error opening %s for reading!", inputFileName)
    panic(errorDesc)
  }

  bufferedInput := bufio.NewReader(inputFile)

  var currentLine []byte

  currentLine, isTruncated, readError := bufferedInput.ReadLine()

  if readError != nil {
    fmt.Fprintf(os.Stdout, "Error reading from %s!: ", inputFileName, readError)
  }

  //First line of file is always the number of test cases...
  numCases,_ := strconv.Atoi(string(currentLine))

  //For each line in the file
  for caseNumber := 1; caseNumber <= numCases && readError == nil; caseNumber++ {
    currentLine, isTruncated, readError = bufferedInput.ReadLine()

    if isTruncated {
      fmt.Fprintf(os.Stdout, "Buffer was too small! Line was truncated to %s ", string(currentLine))
    } else {
      bounds := strings.Fields(string(currentLine))
      lowerBound,_ := strconv.Atoi(bounds[0])
      upperBound,_ := strconv.Atoi(bounds[1])
      countRecycledNumbers(caseNumber, lowerBound, upperBound)
    }

  }

  defer inputFile.Close()
}

func countRecycledNumbers(caseNumber int, lowerBound int, upperBound int) {

  var pairsByNumber = map[int]([]RecycledNumPair){}

  numDistinctPairs := 0

  for number := lowerBound; number <= upperBound; number++ {
    numberAsString := strconv.Itoa(number)
    
    recycled := numberAsString

    for iteration := 1; iteration <= len(numberAsString) - 1; iteration++ {
      recycled = string(recycled[len(recycled)-1]) + recycled[:len(recycled)-1]

      recycledNum,_ := strconv.Atoi(recycled)
      if recycledNum > number && lowerBound <= recycledNum && recycledNum <= upperBound {
        pair := RecycledNumPair{number, recycledNum}
        if !alreadySeen(pairsByNumber, pair) {
          numDistinctPairs++
          pairsByNumber[pair.number] = append(pairsByNumber[pair.number], pair)
        }
      }
    }
  } 

  fmt.Fprintf(os.Stdout, "Case #%d: %d\n", caseNumber, numDistinctPairs)
}

func alreadySeen(pairsByNumber map[int]([]RecycledNumPair), pairToFind RecycledNumPair) bool {
  seen := false

  pairs, present := pairsByNumber[pairToFind.number]

  if !present {
    return false
  }
  
  for _,pair := range pairs {
    if pair.number == pairToFind.number && pair.recycled == pairToFind.recycled {
      seen = true
      break
    }  
  }  

  return seen
}
