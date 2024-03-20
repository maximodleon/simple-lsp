package main

import (
	"bufio"
	"educationallsp/lsp"
	"educationallsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
  logger := getLogger("/Users/yuli/workspace/edulsp/log.txt")
  logger.Println("Hey, I started")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Bytes()
    method, contents, error := rpc.DecodeMessage(msg)

    if error != nil {
      logger.Printf("Got an error: %s", error)
      continue;
    }

    handleMessag(logger, method, contents)
  }
}

func handleMessag(logger *log.Logger, method string, contents []byte) {
  logger.Printf("Received msg with method: %s", method)

  switch method {
   case "initialize":
      var request lsp.InitializeRequest
      if err := json.Unmarshal(contents, &request); err != nil {
        logger.Printf("hey, we could not parse this: %s", err)
      }

    logger.Printf("Connected to: %s %s",
      request.Params.ClientInfo.Name,
      request.Params.ClientInfo.Version)
      msg := lsp.NewInitializedResponse(request.ID)
      reply := rpc.EncodeMEssage(msg)

      writer := os.Stdout
      writer.Write([]byte(reply))
      logger.Println("Reply sent")
  }
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

  if err != nil {
    panic("hey, you did not give me a good file!")
  }

  return log.New(logfile, "[educationallsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
