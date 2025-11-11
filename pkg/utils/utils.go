package utils

import "github.com/google/uuid"

func UUID() string {
	u4 := uuid.New()
	return u4.String()
}
