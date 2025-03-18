package error

import "errors"

var (
	AlreadyExists       = errors.New("already exists")
	NotFound            = errors.New("not found")
	Unauthorized        = errors.New("unauthorized")
	InternalServerError = errors.New("internal server error")
	BadRequest          = errors.New("bad request")
	EmptyRequestBody    = errors.New("empty request body")
	Unauthenticated     = errors.New("unauthenticated")
	ConversionError     = errors.New("conversion error")
	JsonParseError      = errors.New("json parse error")
	CreateError         = errors.New("creation error")
	DeleteError         = errors.New("deletion error")
	UpdateError         = errors.New("update error")
	VerifyError         = errors.New("verification error")
	InvalidInput        = errors.New("invalid input")
)
