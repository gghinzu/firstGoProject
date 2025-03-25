package error

import "errors"

type Result struct {
	Result  string `json:"result"`
	Message error  `json:"message"`
}

// TODO: general response in server (API response object -result data error) 2 const success/error
func NewErrorResult(result string, message error) Result {
	return Result{
		Result:  result,
		Message: message,
	}
}

var (
	AlreadyExists            = NewErrorResult("entity_already_exists", errors.New("already exists"))
	NotFound                 = NewErrorResult("resource_not_found", errors.New("not found"))
	Unauthorized             = NewErrorResult("permission_denied", errors.New("unauthorized"))
	InternalServerError      = NewErrorResult("internal_server_error", errors.New("internal server error"))
	BadRequest               = NewErrorResult("invalid_request", errors.New("bad request"))
	EmptyRequestBody         = NewErrorResult("empty_request_body", errors.New("empty request body"))
	Unauthenticated          = NewErrorResult("authentication_required", errors.New("unauthenticated"))
	ConversionError          = NewErrorResult("data_conversion_failed", errors.New("conversion error"))
	JsonParseError           = NewErrorResult("json_parsing_failed", errors.New("json parse error"))
	CreateError              = NewErrorResult("resource_creation_failed", errors.New("creation error"))
	DeleteError              = NewErrorResult("resource_deletion_failed", errors.New("deletion error"))
	UpdateError              = NewErrorResult("resource_update_failed", errors.New("update error"))
	VerifyError              = NewErrorResult("verification_failed", errors.New("verification error"))
	InvalidInput             = NewErrorResult("invalid_input", errors.New("invalid input"))
	InvalidVerificationCode  = NewErrorResult("invalid_verification_code", errors.New("invalid verification code"))
	ExpiredVerificationCode  = NewErrorResult("verification_code_expired", errors.New("expired verification code"))
	VerificationCodeNotFound = NewErrorResult("verification_code_not_found", errors.New("verification code not found"))
)
