package bot

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/gorilla/websocket"
	"github.com/ipfs/go-cid"
	carv2 "github.com/ipld/go-car/v2"
)

var (
	wsURL = "wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos"
	terms = []string{"#1123", "#1124"}
)

type RepoCommitEvent struct {
	Repo   string      `cbor:"repo"`
	Rev    string      `cbor:"rev"`
	Seq    int64       `cbor:"seq"`
	Since  string      `cbor:"since"`
	Time   string      `cbor:"time"`
	TooBig bool        `cbor:"tooBig"`
	Prev   interface{} `cbor:"prev"`
	Rebase bool        `cbor:"rebase"`
	Blocks []byte      `cbor:"blocks"`

	Ops []RepoOperation `cbor:"ops"`
}

type RepoOperation struct {
	Action string      `cbor:"action"`
	Path   string      `cbor:"path"`
	Reply  []byte      `cbor:"reply"`
	Text   []byte      `cbor:"text"`
	CID    interface{} `cbor:"cid"`
}

func Websocket() {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		slog.Error("Failed to connect to WebSocket", "error", err)
		return
	}
	defer conn.Close()

	slog.Info("Connected to WebSocket", "url", wsURL)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error reading message from WebSocket", "error", err)
			return
		}

		decoder := cbor.NewDecoder(bytes.NewReader(message))

		for {
			var evt RepoCommitEvent
			err := decoder.Decode(&evt)
			if err == io.EOF {
				break
			}
			if err != nil {
				slog.Error("Error decoding CBOR message", "error", err)
				break
			}
			handleEvent(evt)
		}
	}
}

func handleEvent(evt RepoCommitEvent) {
	for _, op := range evt.Ops {
		if op.Action == "create" {
			if len(evt.Blocks) > 0 {
				handleCARBlocks(evt.Blocks, op)
			}
		}
	}
}

func handleCARBlocks(blocks []byte, op RepoOperation) {
	if len(blocks) == 0 {
		return
	}

	reader, err := carv2.NewBlockReader(bytes.NewReader(blocks))
	if err != nil {
		slog.Error("Error creating CAR block reader", "error", err)
		return
	}

	for {
		block, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("Error reading CAR block", "error", err)
			break
		}

		if opTag, ok := op.CID.(cbor.Tag); ok {
			if cidBytes, ok := opTag.Content.([]byte); ok {
				c, err := decodeCID(cidBytes)
				if err != nil {
					slog.Error("Error decoding CID from bytes", "error", err)
					return
				}

				if block.Cid().Equals(c) {
					var decodedData map[string]interface{}
					err := cbor.Unmarshal(block.RawData(), &decodedData)
					if err != nil {
						slog.Error("Error decoding CBOR block", "error", err)
						return
					}

					text, ok := decodedData["text"].(string)
					reply, _ := decodedData["reply"].(map[string]interface{})

					if !ok && reply == nil {
						continue
					}

					if text != "" && containsTerm(text) {
						repost()
					}
				}
			}
		}
	}
}

func containsTerm(text string) bool {
	for _, term := range terms {
		if strings.Contains(strings.ToLower(text), strings.ToLower(term)) {
			return true
		}
	}
	return false
}

func decodeCID(cidBytes []byte) (cid.Cid, error) {
	var c cid.Cid
	c, err := cid.Decode(string(cidBytes))
	if err != nil {
		return c, fmt.Errorf("error decoding CID: %w", err)
	}
	return c, nil
}
