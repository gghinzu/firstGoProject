package error

import "firstGoProject/internal/server"

type Result struct {
	Code    string
	Message string
}

func NewErrorResult(code string, message string) server.Response {
	return server.ErrorResponse(code, message)
}

var (
	AlreadyExists            = NewErrorResult("entity_already_exists", "already exists")
	NotFound                 = NewErrorResult("resource_not_found", "not found")
	Unauthorized             = NewErrorResult("permission_denied", "unauthorized")
	InternalServerError      = NewErrorResult("internal_server_error", "internal server error")
	BadRequest               = NewErrorResult("invalid_request", "bad request")
	EmptyRequestBody         = NewErrorResult("empty_request_body", "empty request body")
	Unauthenticated          = NewErrorResult("authentication_required", "unauthenticated")
	JsonParseError           = NewErrorResult("json_parsing_failed", "json parse error")
	DeleteError              = NewErrorResult("resource_deletion_failed", "deletion error")
	UpdateError              = NewErrorResult("resource_update_failed", "update error")
	VerifyError              = NewErrorResult("verification_failed", "verification error")
	InvalidInput             = NewErrorResult("invalid_input", "invalid input")
	InvalidVerificationCode  = NewErrorResult("invalid_verification_code", "invalid verification code")
	ExpiredVerificationCode  = NewErrorResult("verification_code_expired", "expired verification code")
	VerificationCodeNotFound = NewErrorResult("verification_code_not_found", "verification code not found")
)
