package tag

import (
	"fmt"
	"time"
	"weavelab.xyz/monorail/shared/wlib/uuid"

	"weavelab.xyz/monorail/shared/wlib/werror"
)

type tagType int

const (
	StringType = iota
	IntType
	DurationType
	FloatType
	BoolType
	WErrorType
)

type Tag struct {
	Key       string
	Type      tagType
	StringVal string
	IntVal    int64
	DVal      time.Duration
	FloatVal  float64
	BoolVal   bool
	WErrorVal *werror.Error
}

func String(key string, val string) Tag {
	return Tag{Key: key, Type: StringType, StringVal: val}
}

func Int(key string, val int) Tag {
	return Tag{Key: key, Type: IntType, IntVal: int64(val)}
}

func Int64(key string, val int64) Tag {
	return Tag{Key: key, Type: IntType, IntVal: val}
}

func Duration(key string, val time.Duration) Tag {
	return Tag{Key: key, Type: DurationType, DVal: val}
}

func Float(key string, val float64) Tag {
	return Tag{Key: key, Type: FloatType, FloatVal: val}
}

func Bool(key string, val bool) Tag {
	return Tag{Key: key, Type: BoolType, BoolVal: val}
}

func UUID(key string, val uuid.UUID) Tag {
	return Tag{Key: key, Type: StringType, StringVal: val.String()}
}

func WError(key string, val *werror.Error) Tag {
	return Tag{Key: key, Type: WErrorType, WErrorVal: val}
}

func (t Tag) String() string {
	switch t.Type {
	case StringType:
		return fmt.Sprintf("%s=[%s]", t.Key, t.StringVal)
	case IntType:
		return fmt.Sprintf("%s=[%d]", t.Key, t.IntVal)
	case DurationType:
		return fmt.Sprintf("%s=[%s]", t.Key, t.DVal)
	case FloatType:
		return fmt.Sprintf("%s=[%v]", t.Key, t.FloatVal)
	case BoolType:
		return fmt.Sprintf("%s=[%v]", t.Key, t.BoolVal)
	case WErrorType:
		return fmt.Sprintf("%s=[%v]", t.Key, t.WErrorVal.Error())
	}
	return fmt.Sprint("unknown tag!", t.Key, t.Type)
}
