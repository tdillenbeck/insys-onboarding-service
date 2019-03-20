package wapp

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

//appInfo - this gets set at compile time https://golang.org/cmd/link/ using the -X flag
var (
	appInfo     string // base64 gzipped app info, set at compile time
	appInfoJSON string
)

type appInfoError struct {
	RawValue string
	Error    error
}

// take the base64 gzip encoded app info and decode it
func init() {
	err := appInfoDecode(appInfo)
	if err != nil {
		// log the error
		wlog.WError(werror.Wrap(err))
		appInfoDecodeError(err)
		return
	}
}

// takes a base64 encoded gzipped string and decodes it
func appInfoDecode(s string) error {
	if len(s) == 0 {
		return nil
	}

	buff := bytes.NewBufferString(s)

	g := base64.NewDecoder(base64.StdEncoding, buff)

	j, err := gzip.NewReader(g)
	if err != nil {
		return werror.Wrap(err)
	}

	out, err := ioutil.ReadAll(j)
	if err != nil {
		return werror.Wrap(err)
	}

	appInfoJSON = string(out)

	return nil
}

// set error output if we couldn't decode the value
func appInfoDecodeError(err error) {
	ai := appInfoError{
		RawValue: appInfo,
		Error:    err,
	}

	out, _ := json.Marshal(ai)
	appInfoJSON = string(out)
}

// AppInfo is for retrieving info about an app
func (s statusHandler) AppInfo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(appInfoJSON))
}
