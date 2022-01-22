// Content managed by Project Forge, see [projectforge.md] for details.
package util

import (
	"github.com/google/uuid"
)

func UUIDFromString(s string) *uuid.UUID {
	var retID *uuid.UUID

	if len(s) > 0 {
		s, err := uuid.Parse(s)
		if err == nil {
			retID = &s
		}
	}

	return retID
}

func UUID() uuid.UUID {
	return uuid.New()
}

func UUIDP() *uuid.UUID {
	ret := UUID()
	return &ret
}
