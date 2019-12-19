package zapier

import (
	"bytes"
	"context"
	"encoding/json"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
)

// type ZapierClient interface {
// 	Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error
// }
type ZapierClient struct {
	url string
}

type FirstLoginEventPayload struct {
	Username                string `json:"username"`
	LocationID              string `json:"location_id"`
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
}

func New(url string) *ZapierClient {
	return &ZapierClient{url: url}
}

func (zc *ZapierClient) Send(ctx context.Context, username, locationID, salesforceOpportunityID string) error {
	contentType := "application/json"
	payload := FirstLoginEventPayload{
		Username:                username,
		LocationID:              locationID,
		SalesforceOpportunityID: salesforceOpportunityID,
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
