package mm

import (
  "io"
)

type MpegExtractor struct {}

func (ex MpegExtractor) Extract(rs io.ReadSeeker) (Metadata, error) {
  return Metadata{}, nil
}
