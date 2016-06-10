package mm

import (
  "io"
  "os"
  flac "github.com/nkcmr/mm-flac"
  mpeg "github.com/nkcmr/mm-mpeg"
)

type nilExtractor struct {}

// Parse takes a filename and automatically looks up an extractor to use
func Parse (rs io.ReadSeeker) (Metadata, error) {
  extractor, err := initExtractor(rs)
  if err != nil {
    return Metadata{}, err
  }
  return ParseWithExtractor(rs, extractor)
}

// ParseWithExtractor will extract music metadata out of a file with a specific
// extractor
func ParseWithExtractor (rs io.ReadSeeker, extractor Extractor) (Metadata, error) {
  return Metadata{}, nil
}

func pickExtractor (rs io.ReadSeeker) (string, error) {
  extractors := map[string]func (r io.ReadSeeker) (bool, error){
    "flac": flac.CanExtract,
    "mpeg": mpeg.CanExtract,
  }
  for _type, canExtract := range extractors {
    _, err := rs.Seek(0, os.SEEK_SET)
    if err != nil {
      return "", err
    }
    ok, err := canExtract(rs)
    if err != nil {
      return "", err
    }
    if ok {
      return _type, nil
    }
  }
  return "", nil
}

func initExtractor (rs io.ReadSeeker) (Extractor, error) {
  ex, err := pickExtractor(rs)
  if err != nil {
    return nilExtractor{}, err
  }
  switch ex {
  case "flac":
    return flac.NewExtractor(), nil
  case "mpeg":
    return mpeg.NewExtractor(), nil
  default:
    return nilExtractor{}, &MusicMetadataError{E_UNSUPPORTED_TYPE, "unsupported audio type"}
  }
}
