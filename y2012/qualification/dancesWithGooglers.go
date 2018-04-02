package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
  "math"
  "flag"
  "sort"
)

var inputFileName *string = flag.String("file", "", "The input file for this script")

type Triplet struct {
  judge1,judge2,judge3 int
}

func (t Triplet) totalScore() int {
  return t.judge1 + t.judge2 + t.judge3
}

func (t Triplet) isValid() bool {
  isValidTriplet := int(math.Abs(float64(t.judge1 - t.judge2))) <= 2 &&
                    int(math.Abs(float64(t.judge1 - t.judge3))) <= 2 &&
                    int(math.Abs(float64(t.judge2 - t.judge3))) <= 2

  return isValidTriplet
}

func (t Triplet) bestScore() int {
  tripletSlice := []int{t.judge1, t.judge2, t.judge3}
  sort.Ints(tripletSlice)
  return tripletSlice[len(tripletSlice) - 1]
}

func (t Triplet) isSurprising() bool {
  isSurprisingTriplet := int(math.Abs(float64(t.judge1 - t.judge2))) == 2 ||
                         int(math.Abs(float64(t.judge1 - t.judge3))) == 2 ||
                         int(math.Abs(float64(t.judge2 - t.judge3))) == 2

  return isSurprisingTriplet
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
      findNumGooglers(caseNumber, strings.Fields(string(currentLine)))
    }

  }

  defer inputFile.Close()
}

func findNumGooglers(caseNumber int, input []string) {

  numGooglers,_ := strconv.Atoi(input[0])
  numSurprisingTriplets,_ := strconv.Atoi(input[1])
  minBestScore,_ := strconv.Atoi(input[2])

  var scores []int
  startingIndex := 3
  var score int

  for offset := 0 ; offset < numGooglers; offset++ {
    score,_ = strconv.Atoi(input[startingIndex + offset])
    scores = append(scores, score)
  }

  //Now we have all the scores let's generate our mappings from score to
  //num Googlers and from score to candidate triplets...
  var scoreToNumGooglers = map[int]int{}
  var scoreToTriplets = map[int]([]Triplet){}

  for _,totalScore := range scores {
    scoreToNumGooglers[totalScore] += 1
  }

  maxScore := 10

  //Triplet time!
  for firstScore := 0; firstScore <= maxScore; firstScore++ {
    for secondScore := 0; secondScore <= maxScore; secondScore++ {
      for thirdScore := 0; thirdScore <= maxScore; thirdScore++ {
        triplet := Triplet{firstScore, secondScore, thirdScore}
        if triplet.isValid() {
          tripletScore := triplet.totalScore()
          _,scoreRepresented := scoreToNumGooglers[tripletScore]

          if scoreRepresented && triplet.bestScore() >= minBestScore {
            scoreToTriplets[tripletScore] = append(scoreToTriplets[tripletScore], triplet)
          }
        }
      }
    }
  }

  maxGooglers := 0
  surprisesUsed := 0

  for totalScore,numGooglers := range scoreToNumGooglers {
    for _,triplet := range scoreToTriplets[totalScore] {
      tripletFound := false
      switch {
        case triplet.isSurprising() && !hasNonSurprisingTriplet(scoreToTriplets[totalScore]) && surprisesUsed < numSurprisingTriplets:
          tripletFound = true
          var googlersWithSurprises int
          surprisesAvailable := numSurprisingTriplets - surprisesUsed
          googlersWithSurprises += int(math.Min(float64(numGooglers), float64(surprisesAvailable)))

          maxGooglers += googlersWithSurprises
          surprisesUsed += googlersWithSurprises
        case !triplet.isSurprising():
          tripletFound = true
          maxGooglers += numGooglers
      }

      if tripletFound {
        break
      }
    }
  }

  fmt.Fprintf(os.Stdout, "Case #%d: %d\n", caseNumber, maxGooglers)
}

func hasNonSurprisingTriplet(triplets []Triplet) bool {
  for _,triplet := range triplets {
    if !triplet.isSurprising() {
      return true
    }
  }
  return false
}
