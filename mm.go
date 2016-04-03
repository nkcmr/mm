package mm

import (
  "io"
  "os"
  flac "github.com/nkcmr/mm-flac"
  mpeg "github.com/nkcmr/mm-mpeg"
)

type nilExtractor struct {}

func Parse (path string) (Metadata, error) {
  extractor, err := initExtractor(path)
  if err != nil {
    return Metadata{}, err
  }
  return ParseWithExtractor(path, extractor)
}

func ParseWithExtractor (path string, extractor Extractor) (Metadata, error) {
  return Metadata{}, nil
}

func pickExtractor (path string) (string, error) {
  extractors := map[string]func (r io.ReadSeeker) (bool, error){
    "flac": flac.CanExtract,
    "mpeg": mpeg.CanExtract,
  }
  file, err := os.Open(path)
  if err != nil {
    return "", err
  }
  for _type, canExtract := range extractors {
    ok, err := canExtract(file)
    if err != nil {
      return "", err
    }
    if ok {
      return _type, nil
    }
  }
  return "", nil
}

func initExtractor (path string) (Extractor, error) {
  ex, err := pickExtractor(path)
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
