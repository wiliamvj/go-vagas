package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type CreateRecordProps struct {
	DIDResponse *DIDResponse
	Resource    string
	URI         string
	CID         string
}

func createRecord(r *CreateRecordProps) error {
	body := map[string]interface{}{
		"$type":      r.Resource,
		"collection": r.Resource,
		"repo":       r.DIDResponse.DID,
		"record": map[string]interface{}{
			"subject": map[string]interface{}{
				"uri": r.URI,
				"cid": r.CID,
			},
			"createdAt": time.Now(),
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		slog.Error("Error marshalling request", "error", err, "resource", r.Resource)
		return err
	}

	url := fmt.Sprintf("%s/com.atproto.repo.createRecord", API_URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		slog.Error("Error creating request", "error", err, "r.Resource", r.Resource)
		return nil
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.DIDResponse.AccessJwt))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request", "error", err, "r.Resource", r.Resource)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		slog.Error("Unexpected status code", "status", resp, "r.Resource", r.Resource)
		return nil
	}

	slog.Info("Published successfully", "resource", r.Resource)

	return nil
}
