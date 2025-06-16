package gelbooru

import "fmt"

type GelbooruError struct {
	Code int
}

func (e GelbooruError) Error() string {
	return fmt.Sprintf("gelbooru request error: code %d", e.Code)
}
