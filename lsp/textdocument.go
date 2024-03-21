package lsp

type TextDocumentItem struct {
  /**
  * The text document's URI
  */
  URI string `json:"uri"`

  /**
  * The text document's language identifier
  */
  LanguageID string `json:"languageId"`

  /**
  * The number of this document
  */
  Version int `json:"version"`;

  /**
  * The content of the document
  */
  Text string `json:"text"`
}

type TextDocumentPositionParams struct {
   TextDocument TextDocumentIdentifier `json:"textDocument"`
   Position Position `json:"position"`
}

type Position struct {
   Line int `json:"line"`
   Character int `json:"character"`
}

type TextDocumentIdentifier struct {
    URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
   TextDocumentIdentifier
   Version int `json:"version"`
}

/**
* Did open
*/
type DidOpenTextDocumentNotification struct {
    Notification
    Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
    TextDocument TextDocumentItem `json:"textDocument"`;
}

/**
* Did change
*/
type TextDocumentDidChangeNotification struct {
   Notification
   Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
    TextDocument VersionTextDocumentIdentifier `json:"textDocument"`
    ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
    Text string `json:"text"`
}

/**
* Hover
*/

type HoverRequest struct {
  Request
  Params HoverParams `json:"params"`;
}

type HoverParams struct {
  TextDocumentPositionParams
}

type HoverResponse struct {
  Response
  Result HoverResult `json:"result"`
}

type HoverResult struct {
  Contents string `json:"contents"`
}
