package helper

import (
	"errors"

	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/google/uuid"
)

var ErrCanNotGenerateUUID = errors.New("can not generate uuid")

func UUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Default().Panic(ErrCanNotGenerateUUID.Error())
	}

	return id.String()
}

func UUIDMask() string {
	return "\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b\n"
}
