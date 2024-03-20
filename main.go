package main

import (
	"bufio"
	"educationallsp/analysis"
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

  state := analysis.NewState()

  for scanner.Scan() {
    msg := scanner.Bytes()
    method, contents, error := rpc.DecodeMessage(msg)

    if error != nil {
      logger.Printf("Got an error: %s", error)
      continue;
    }

    handleMessag(logger, state, method, contents)
  }
}

func handleMessag(logger *log.Logger, state analysis.State, method string, contents []byte) {
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
  case "textDocument/didOpen":
      var request lsp.DidOpenTextDocumentNotification
      if err := json.Unmarshal(contents, &request); err != nil {
        logger.Printf("textDocument/didOpen: %s", err)
        return
      }

      logger.Printf("Connected to: %s", request.Params.TextDocument.URI)
      state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
  case "textDocument/didChange":
      var request lsp.TextDocumentDidChangeNotification
      if err := json.Unmarshal(contents, &request); err != nil {
        logger.Printf("textDocument/didChange: %s", err)
        return
      }

      logger.Printf("Changed: %s", request.Params.TextDocument.URI)

       for _, change := range request.Params.ContentChanges {
          state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
       }
  }
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

  if err != nil {
    panic("hey, you did not give me a good file!")
  }

  return log.New(logfile, "[educationallsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
