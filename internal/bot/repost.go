package bot

import (
	"log/slog"
)

func repost(p *Post) error {
	token, err := getToken()
	if err != nil {
		slog.Error("Error getting token", "error", err)
		return err
	}

	resource := &CreateRecordProps{
		Post:        p,
		DIDResponse: token,
		Resource:    "app.bsky.feed.repost",
		URI:         p.Reply.Root.Uri,
		CID:         p.Reply.Root.Cid,
	}

	err = createRecord(resource)
	if err != nil {
		slog.Error("Error creating record", "error", err, "resource", resource.Resource)
		return err
	}

	resource.Resource = "app.bsky.feed.like"
	err = createRecord(resource)
	if err != nil {
		slog.Error("Error creating record", "error", err, "resource", resource.Resource)
		return err
	}

	return nil
}
