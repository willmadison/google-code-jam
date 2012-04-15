package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "flag"
)

var inputFileName *string = flag.String("file", "", "The input file for this script")

var dictionary = map[string]string {
  "a":"y",  
  "b":"h",  
  "c":"e",  
  "d":"s",  
  "e":"o",  
  "f":"c",  
  "g":"v",  
  "h":"x",  
  "i":"d",  
  "j":"u",  
  "k":"i",  
  "l":"g",  
  "m":"l",  
  "n":"b",  
  "o":"k",  
  "p":"r",  
  "q":"z",  
  "r":"t",  
  "s":"n",  
  "t":"w",  
  "u":"j",  
  "v":"p",  
  "w":"f",  
  "x":"m",  
  "y":"a",  
  "z":"q",  
  " ":" ",  
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
      //fmt.Println(string(currentLine))  
      translate(caseNumber, string(currentLine))
    }

  }

  defer inputFile.Close()
}

func translate(caseNumber int, currentLine string) {

  result := ""
  
  for index := 0; index < len(currentLine); index++ {
    
    result += dictionary[string(currentLine[index])]
  } 

  fmt.Fprintf(os.Stdout, "Case #%d: %s\n", caseNumber, result)
}

