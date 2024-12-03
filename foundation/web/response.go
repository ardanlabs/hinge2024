package web

import (
	"context"
	"fmt"
	"net/http"
)

type httpStatus interface {
	HTTPStatus() int
}

func Respond(ctx context.Context, w http.ResponseWriter, dataModel Encoder) error {
	data, contentType, err := dataModel.Encode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", contentType)

	var statusCode = http.StatusOK

	switch v := dataModel.(type) {
	case httpStatus:
		statusCode = v.HTTPStatus()

	case error:
		statusCode = http.StatusInternalServerError

	default:
		if dataModel == nil {
			statusCode = http.StatusNoContent
		}
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
