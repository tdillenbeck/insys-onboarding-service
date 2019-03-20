package wnsq

import "encoding/json"

func isJSON(b []byte) bool {
	for _, v := range b {
		switch v {
		// should we trim spaces?
		case '{':
			return true
		}
		break
	}
	return false
}

type jsonWrapper struct {
	OpenTracing string `json:",omitempty"`
}

func injectorJSON(md map[string]string, b []byte) ([]byte, error) {

	if b[0] != '{' {
		return b, nil // unable to inject if not an obect
	}

	if md == nil || len(md) == 0 {
		return b, nil
	}

	oo, err := json.Marshal(md)
	if err != nil {
		return nil, err
	}

	w := jsonWrapper{
		OpenTracing: string(oo),
	}

	o, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	d := make([]byte, len(o)+len(b)-1)
	d[0] = '{'
	_ = copy(d[1:len(o)-1], o[1:len(o)-1])
	d[len(o)-1] = ','
	copy(d[len(o):], b[1:])

	return d, nil
}

func extractJSON(b []byte) (map[string]string, error) {
	// we can only unwrap object type JSON messages
	w := jsonWrapper{}
	err := json.Unmarshal(b, &w)
	if err != nil {
		return nil, err
	}

	md := make(map[string]string)
	err = json.Unmarshal([]byte(w.OpenTracing), &md)
	if err != nil {
		return nil, err
	}

	return md, nil
}
