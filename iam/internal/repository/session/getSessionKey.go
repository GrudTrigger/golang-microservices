package session

import (
	"fmt"
)

func (r *repository) GetSessionKey(newUuid string) string {
	return fmt.Sprintf("%s%s", cacheKeyPrefix, newUuid)
}
