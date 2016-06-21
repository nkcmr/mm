package mm

import (
  "io"
  "os"
)

// Parse takes a filename and automatically looks up an extractor to use
func Parse (rs io.ReadSeeker) (*Metadata, error) {
  extractor, err := initExtractor(rs)
  if err != nil {
    return new(Metadata), err
  }
  return ParseWithExtractor(rs, extractor)
}

// ParseWithExtractor will extract music metadata out of a file with a specific
// extractor
func ParseWithExtractor (rs io.ReadSeeker, ex Extractor) (*Metadata, error) {
  if _, err := rs.Seek(0, os.SEEK_SET); err != nil {
    return new(Metadata), err
  }
  return ex.Extract(rs)
}

func initExtractor (rs io.ReadSeeker) (Extractor, error) {
  var ex Extractor
  var err error
  var choice string
  choice, err = DetectAudioType(rs)
  if err != nil {
    return ex, err
  }
  switch choice {
  case "flac":
    ex = FlacExtractor{}
  case "mpeg":
    ex = MpegExtractor{}
  default:
    err = &MusicMetadataError{ErrorUnsupportedType, "unsupported audio type"}
  }
  return ex, err
}
