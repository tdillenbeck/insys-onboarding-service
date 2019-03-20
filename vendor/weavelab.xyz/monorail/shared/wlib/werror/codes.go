package werror

import "net/http"

type Code int
type Class string

const (
	CodeOK = Code(0)

	// These error codes should map one to one with the gRPC error codes
	CodeCanceled           = Code(1)
	CodeUnknown            = Code(2)
	CodeInvalidArgument    = Code(3)
	CodeDeadlineExceeded   = Code(4)
	CodeNotFound           = Code(5)
	CodeAlreadyExists      = Code(6)
	CodePermissionDenied   = Code(7)
	CodeResourceExhausted  = Code(8)
	CodeFailedPrecondition = Code(9)
	CodeAborted            = Code(10)
	CodeOutOfRange         = Code(11)
	CodeUnimplemented      = Code(12)
	CodeInternal           = Code(13)
	CodeUnavailable        = Code(14)
	CodeDataLoss           = Code(15)
	CodeUnauthenticated    = Code(16)
	CodeConflict           = Code(17) // TODO: this isn't a grpc error code and may cause _conflicts_ in the future

	Unknown Class = "0xx"
	// Success represents outcomes that achieved the desired results.
	Success Class = "2xx"
	// ClientError represents errors that were the client's fault.
	ClientError Class = "4xx"
	// ServerError represents errors that were the server's fault.
	ServerError Class = "5xx"
)

var codeMap = map[Code]int{
	CodeOK:                 http.StatusOK,
	CodeCanceled:           http.StatusRequestTimeout,
	CodeUnknown:            http.StatusInternalServerError,
	CodeInvalidArgument:    http.StatusBadRequest,
	CodeDeadlineExceeded:   http.StatusRequestTimeout,
	CodeNotFound:           http.StatusNotFound,
	CodeAlreadyExists:      http.StatusUnprocessableEntity,
	CodePermissionDenied:   http.StatusUnauthorized,
	CodeResourceExhausted:  http.StatusTooManyRequests,
	CodeFailedPrecondition: http.StatusPreconditionFailed,
	CodeAborted:            http.StatusInternalServerError,
	CodeOutOfRange:         http.StatusBadRequest,
	CodeUnimplemented:      http.StatusNotImplemented,
	CodeInternal:           http.StatusInternalServerError,
	CodeUnavailable:        http.StatusServiceUnavailable,
	CodeDataLoss:           http.StatusInternalServerError,
	CodeUnauthenticated:    http.StatusUnauthorized,
	CodeConflict:           http.StatusConflict,
}

// HttpCode returns the related http status code based on our predefined codes
func HttpCode(code Code) int {
	httpCode, ok := codeMap[code]
	if ok {
		return httpCode
	}

	return http.StatusInternalServerError
}

// ErrorClass returns the class of the given error
func ErrorClass(code Code) Class {

	switch code {
	// Success or "success"
	case CodeOK:
		return Success

	// Client errors
	case CodeCanceled, CodeInvalidArgument, CodeNotFound, CodeAlreadyExists,
		CodePermissionDenied, CodeUnauthenticated, CodeFailedPrecondition,
		CodeOutOfRange, CodeConflict:
		return ClientError

	// Server errors
	case CodeDeadlineExceeded, CodeResourceExhausted, CodeAborted,
		CodeUnimplemented, CodeInternal, CodeUnavailable, CodeDataLoss:
		return ServerError

	// Not sure
	case CodeUnknown:
		return Unknown
	default:
		return Unknown
	}

}

func HttpToWeaveCode(i int) Code {
	for k, v := range codeMap {
		if v == i {
			return k
		}
	}
	return CodeUnknown
}
