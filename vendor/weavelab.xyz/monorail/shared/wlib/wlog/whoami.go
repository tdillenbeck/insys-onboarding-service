package wlog

import (
	"os"
	"path/filepath"

	"weavelab.xyz/monorail/shared/wlib/version"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlogd/wlogproto"
)

const magicWhoAmI = "WłøǥđWħøȺmɨ" // this is funky unicode to prevent anyone from sending it by accident

var whoAmILog = wlogproto.Log{
	Level:        wlogproto.Level_INFO,
	Message:      magicWhoAmI,
	TagsString:   make(map[string]string),
	TagsInt:      make(map[string]int32),
	TagsInt64:    make(map[string]int64),
	TagsFloat:    make(map[string]float32),
	TagsBool:     make(map[string]bool),
	TagsDuration: make(map[string]*wlogproto.Duration),
}

func AddWhoAmI(t tag.Tag) {
	switch t.Type {
	case tag.StringType:
		whoAmILog.TagsString[t.Key] = t.StringVal
	case tag.IntType:
		whoAmILog.TagsInt64[t.Key] = t.IntVal
	case tag.DurationType:
		whoAmILog.TagsDuration[t.Key] = &wlogproto.Duration{Duration: t.IntVal}
	case tag.FloatType:
		whoAmILog.TagsFloat[t.Key] = float32(t.FloatVal)
	case tag.BoolType:
		whoAmILog.TagsBool[t.Key] = t.BoolVal
	case tag.WErrorType:
		// ignoring this one for whoami
	}
}

func init() {
	info := version.Info()
	// default whoamis from wlib/version

	AddWhoAmI(tag.String("name", filepath.Base(os.Args[0])))
	AddWhoAmI(tag.String("version", info.Version))
	AddWhoAmI(tag.String("goversion", info.GoVersion))
	AddWhoAmI(tag.String("commitHash", info.GitHash))
	AddWhoAmI(tag.String("branch", info.GitBranch))
	AddWhoAmI(tag.String("hostname", info.Hostname))
	AddWhoAmI(tag.Int64("modTime", info.FileModificationTime.Unix()))
}
