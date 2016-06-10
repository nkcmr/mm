package mm

import (
  "io"
  "os"
  "reflect"
)

type audioFileSignature struct {
  Offset int64
  Markers [][]byte
}

var signatures = map[string]audioFileSignature{
  "flac": audioFileSignature{Offset: 0, Markers: [][]byte{[]byte{'f', 'L', 'a', 'C'}}},
  "mpeg": audioFileSignature{Offset: 0, Markers: [][]byte{[]byte{'I', 'D', '3', 2}, []byte{'I', 'D', '3', 3}, []byte{'I', 'D', '3', 4}}},
}

func DetectorAudioType(rs io.ReadSeeker) (string, error) {
  for kind, sig := range signatures {
    for _, marker := range sig.Markers {
      if _, err := rs.Seek(sig.Offset, os.SEEK_SET); err != nil {
        return "", err
      }
      data := make([]byte, len(marker))
      _, err := io.ReadFull(rs, data)
      if err != nil {
        return "", err
      }
      if reflect.DeepEqual(marker, data) {
        return kind, nil
      }
    }
  }
  return "", &MusicMetadataError{E_UNSUPPORTED_TYPE, "unsupported audio type"}
}
