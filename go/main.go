package main

import (
  log "github.com/sirupsen/logrus"
  "strconv"
)

func main() {
  log.SetFormatter(&log.JSONFormatter{})

  addition(2, 3, false)
  addition(-1, 3, false)
  addition(10, -1, false)
  addition(-1, -1, false)
  addition(5, 7, true)
}

func addition(number1 int, number2 int, errorCondition bool) string {
  var result int = -1

  if (number1 < 0) {
    log.WithFields(log.Fields{
      "number1": number1,
    }).Warn("addition function: number1 is negative integer")
  }

  if (number2 < 0) {
    log.WithFields(log.Fields{
      "number2": number2,
    }).Warn("addition function: number2 is negative integer")
  }

  result = number1 + number2

  if (errorCondition) {
    log.WithFields(log.Fields{
      "number1": number1,
      "number2": number2,
      "result": result,
    }).Fatal("addition function called")
  } else {
    log.WithFields(log.Fields{
      "number1": number1,
      "number2": number2,
      "result": result,
    }).Info("addition function called")
  }

  return strconv.Itoa(result)
}