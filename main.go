package main

import (
	"bufio"
	"educationallsp/rpc"
	"os"
  "log"
)

func main() {
  logger := getLogger("/Users/yuli/workspace/edulsp/log.txt")
  logger.Println("Hey, I started")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Text()
    handleMessag(logger, msg)
  }
}

func handleMessag(logger *log.Logger, msg any) {
  logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

  if err != nil {
    panic("hey, you did not give me a good file!")
  }

  return log.New(logfile, "[educationallsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
