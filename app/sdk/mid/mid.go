package mid

import "github.com/ardanlabs/service/foundation/web"

func isError(e web.Encoder) error {
	err, isError := e.(error)
	if isError {
		return err
	}
	return nil
}
