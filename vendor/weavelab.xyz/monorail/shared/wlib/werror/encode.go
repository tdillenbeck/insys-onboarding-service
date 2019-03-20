package werror

import (
	"encoding/json"
)

type encodedErrorBody struct {
	Cause     string   `json:"cause,omitempty"`
	Messages  []string `json:"messages,omitempty"`
	ClientMsg string   `json:"clientMsg,omitempty"`
	Code      Code     `json:"code,omitempty"`
	Tags      []Tags   `json:"tags,omitempty"`
}

func Encode(err error) ([]byte, error) {
	werr, ok := err.(*Error)
	if !ok {
		werr = Wrap(err)
	}

	cause := ""
	if werr.err != nil {
		cause = werr.err.Error()
	}

	encodedError := encodedErrorBody{
		Cause:     cause,
		Messages:  werr.messages,
		ClientMsg: werr.clientMsg,
		Code:      werr.code,
		Tags:      werr.tags,
	}

	encoded, err := json.Marshal(encodedError)
	if err != nil {
		return nil, Wrap(err, "unable to encode")
	}

	return encoded, nil
}

func Decode(bytes []byte) (*Error, error) {
	var body encodedErrorBody
	err := json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, Wrap(err, "unable to unmarshal error body")
	}

	werr := New(body.Cause)
	werr.clientMsg = body.ClientMsg
	werr.messages = body.Messages
	werr.code = body.Code
	werr.tags = body.Tags

	return werr, nil
}
