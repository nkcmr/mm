package mm

import "fmt"

type ErrorCode int

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
