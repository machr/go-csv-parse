package main

import (
  "fmt"
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
  file, err := os.Open("emails.csv")
  if err != nil {
    log.Fatal(err)
  }

  // close file after program runs
  defer file.Close()

  // read values using csv.Reader
  csvReader := csv.NewReader(file)
  data, err := csvReader.ReadAll()
  if err != nil {
    log.Fatal(err) 
  }

  // convert "record" to array of structs (object)
  userList := createUserList(data)

  fmt.Printf("%+v\n", userList)

}


