package errors

// error response on json:api fromat https://jsonapi.org/

type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Header    string `json:"header,omitempty"`
}

type ErrorLink struct {
	About string `json:"about,omitempty"`
}

type JSONAPIError struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source string `json:"source,omitempty"`
	Meta   any    `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Errors []JSONAPIError `json:"errors"`
}

const (
	CodeValidationError  = "VALIDATION_ERROR"
	CodeNotFound         = "NOT_FOUND"
	CodeUnauthorized     = "UNAUTHORIZED"
	CodeForbidden        = "FORBIDDEN"
	CodeInternalError    = "INTERNAL_ERROR"
	CodeBadRequest       = "BAD_REQUEST"
	CodeConflict         = "CONFLICT"
	CodeTooManyRequests  = "TOO_MANY_REQUESTS"
	CodeMethodNotAllowed = "METHOD_NOT_ALLOWED"
	CodeUnsupportedMedia = "UNSUPPORTED_MEDIA_TYPE"
)

const (
	TitleValidationError  = "Validation Error"
	TitleNotFound         = "Resource Not Found"
	TitleUnauthorized     = "Unauthorized"
	TitleForbidden        = "Forbidden"
	TitleInternalError    = "Internal Server Error"
	TitleBadRequest       = "Bad Request"
	TitleConflict         = "Conflict"
	TitleTooManyRequests  = "Too Many Requests"
	TitleMethodNotAllowed = "Method Not Allowed"
)
