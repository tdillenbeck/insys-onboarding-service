package zapier

import (
	"bytes"
	"context"
	"encoding/json"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
)

type ZapierClient struct {
	url string
}

type FirstLoginEventPayload struct {
	Username   string `json:"username"`
	LocationID string `json:"location_id"`
}

func New(url string) *ZapierClient {
	return &ZapierClient{url: url}
}

func (zc *ZapierClient) Send(ctx context.Context, username, locationID string) error {
	contentType := "application/json"
	payload := FirstLoginEventPayload{
		Username:   username,
		LocationID: locationID,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return werror.Wrap(err, "could not marshal porting data zapier into json")
	}
	r := bytes.NewReader(body)

	_, err = whttpclient.Post(ctx, zc.url, contentType, r)
	if err != nil {
		return werror.Wrap(err, "could not post to zapier")
	}

	return nil
}
