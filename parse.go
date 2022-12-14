package main

import (
  "fmt"
  "io"
  "os"
  "encoding/csv"
  "log"
)

type UserRecord struct {
  FirstName string
  LastName string
  Email string
}


func createUserList(data [][]string) []UserRecord {
  // Create a user list with the shape of UserRecord struct (think object in JS)
  var userList []UserRecord
  
  // i is index, 
  for i, line := range data {
    if i > 0 { // omit header line
      var record UserRecord
      for j, field := range line {
        if j == 0 {
          record.FirstName = field
        } else if j == 1 {
          record.LastName = field
        } else if j == 2 {
          record.Email = field
        }
      }
     userList = append(userList, record)
    }
  }

  return userList
}

func main() {
  // open the file
  f, err := os.Open("emails.csv")
  if err != nil {
    log.Fatal(err)
  }

  // close the file after the program has run
  defer f.Close()

  // read csv values using csv.Reader
  csvReader := csv.NewReader(f)
  for {
    record, err := csvReader.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    // print the user list array - If the value is a struct, the %+v variant will include the struct’s field names.
    fmt.Printf("%+v\n", record)
  }

}


