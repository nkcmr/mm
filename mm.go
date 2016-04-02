package mm

import (
  "fmt"
  "mime"
  "path/filepath"
  "os"
  flac "github.com/nkcmr/mm-flac"
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

func initExtractor (path string) (Extractor, error) {
  file, err := os.Open(path)
  if err != nil {
    return nilExtractor{}, err
  }
  ok, err := flac.CanExtract(file)
  if err != nil {
    return nilExtractor{}, err
  }
  if ok {
    fmt.Println("coool! flac can extract!")
    return nilExtractor{}, nil
  }
  fmt.Println(mime.TypeByExtension(filepath.Ext(path)))
  return nilExtractor{}, nil
}
