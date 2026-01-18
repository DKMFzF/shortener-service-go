package resizeUrl

import (
	"shortener/internal/errors"
)

const (
	ERROR_CODE_URL_EMPTY   = "EMPTY_URL"
	ERROR_CODE_URL_TO_LONG = "URL_TOO_LONG"
)

const (
	ERROR_TITLE_URL_EMPTY   = "URL cannot be empty"
	ERROR_TITLE_UTL_TO_LONG = "URL exceeds maximum length of 2000 characters"
)

func ResizeUrl(url string) (string, error) {
	if url == "" {
		return "", errors.NewServiceError(
			ERROR_CODE_URL_EMPTY,
			ERROR_TITLE_URL_EMPTY,
			nil,
		)
	}

	if len(url) > 2000 {
		return "", errors.NewServiceError(
			ERROR_CODE_URL_TO_LONG,
			ERROR_TITLE_UTL_TO_LONG,
			nil,
		)
	}

	return "short-code", nil
}
