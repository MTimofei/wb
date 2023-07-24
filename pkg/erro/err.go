package erro

import "fmt"

func IsError(msg string, err error) error {
	if err != nil {
		err = fmt.Errorf("%s:%w", msg, err)
		return err
	}
	return nil
}
