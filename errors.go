package mm

import "fmt"

type ErrorCode uint

const (
  ErrorUnsupportedType ErrorCode = iota
)

type MusicMetadataError struct {
  code ErrorCode
  message string
}

func (e *MusicMetadataError) Error () string {
  return fmt.Sprintf("[music metadata error] - %s", e.message)
}
