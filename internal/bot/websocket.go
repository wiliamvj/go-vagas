package bot

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/fxamacker/cbor/v2"
	"github.com/gorilla/websocket"
	"github.com/ipfs/go-cid"
	carv2 "github.com/ipld/go-car/v2"
	"github.com/wiliamvj/go-vagas/internal/utils"
)

var (
	wsURL = "wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos"
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
	Reply  *Reply      `cbor:"reply"`
	Text   []byte      `cbor:"text"`
	CID    interface{} `cbor:"cid"`
}

type Reply struct {
	Parent Parent `json:"parent"`
	Root   Root   `json:"root"`
}

type Parent struct {
	Cid string `json:"cid"`
	Uri string `json:"uri"`
}

type Root struct {
	Cid string `json:"cid"`
	Uri string `json:"uri"`
}

type Post struct {
	Type  string `json:"$type"`
	Text  string `json:"text"`
	Reply *Reply `json:"reply"`
}

func Websocket() error {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		slog.Error("Failed to connect to WebSocket", "error", err)
		return err
	}
	defer conn.Close()

	slog.Info("Connected to WebSocket", "url", wsURL)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error reading message from WebSocket", "error", err)
			continue
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
			err = handleEvent(evt)
			if err != nil {
				return err
			}
		}
	}
}

func handleEvent(evt RepoCommitEvent) error {
	for _, op := range evt.Ops {
		if op.Action == "create" {
			if len(evt.Blocks) > 0 {
				err := handleCARBlocks(evt.Blocks, op)
				if err != nil {
					slog.Error("Error handling CAR blocks", "error", err)
					return err
				}
			}
		}
	}

	return nil
}

func handleCARBlocks(blocks []byte, op RepoOperation) error {
	if len(blocks) == 0 {
		return errors.New("no blocks to process")
	}

	reader, err := carv2.NewBlockReader(bytes.NewReader(blocks))
	if err != nil {
		slog.Error("Error creating CAR block reader", "error", err)
		return err
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
					continue
				}

				if block.Cid().Equals(c) {
					var post Post
					err := cbor.Unmarshal(block.RawData(), &post)
					if err != nil {
						slog.Error("Error decoding CBOR block", "error", err)
						continue
					}

					if post.Text == "" || post.Reply == nil {
						continue
					}

					if utils.FilterTerms(post.Text) {
						repost(&post)
					}
				}
			}
		}
	}

	return nil
}

func decodeCID(cidBytes []byte) (cid.Cid, error) {
	var c cid.Cid
	c, err := cid.Decode(string(cidBytes))
	if err != nil {
		return c, fmt.Errorf("error decoding CID: %w", err)
	}

	return c, nil
}
