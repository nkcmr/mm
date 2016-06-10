package mm

import (
  "io"
)

type Extractor interface {
  Extract(rs io.ReadSeeker) (Metadata, error)
}
