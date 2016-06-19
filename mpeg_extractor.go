package mm

import (
  "io"
  "os"
  "fmt"
  "strings"
)

type id3Flags struct {
  unsynchronisation bool
  extendedHeader bool
  experimental bool
}

func (f id3Flags) String() string {
  return strings.Join([]string{
    fmt.Sprintf("unsynchronisation: %t", f.unsynchronisation),
    fmt.Sprintf("extendedHeader: %t", f.extendedHeader),
    fmt.Sprintf("experimental: %t", f.experimental),
  }, "\n  ")
}

type id3Version struct {
  major uint8
  minor uint8
}

func (v id3Version) String() string {
  return fmt.Sprintf("ID3v%d.%d", v.major, v.minor)
}

type id3Header struct {
  version id3Version
  flags id3Flags
  size uint32
}

func (h id3Header) String() string {
  return strings.Join([]string{
    "id3 header", 
    fmt.Sprintf("version: %s", h.version),
    "flags:",
    fmt.Sprintf("  %s", h.flags),
    fmt.Sprintf("size: %d", h.size),
  }, "\n")
}

func readHeader(data []byte) (h id3Header) {
  h.version = id3Version{major: uint8(data[0]), minor: uint8(data[1])}
  h.flags = id3Flags{}
  flagsByte := uint8(data[2])
  h.flags.unsynchronisation = (flagsByte & 0x80) != 0
  h.flags.extendedHeader = (flagsByte & 0x40) != 0
  h.flags.experimental = (flagsByte & 0x20) != 0
  h.size = uint32(((data[6] & 0xFF) | ((data[5] & 0xFF) << 7) | ((data[4] & 0xFF) << 14) | ((data[3] & 0xFF) << 21)) + 10)
  return h
}

type MpegExtractor struct {}

func (ex MpegExtractor) Extract(rs io.ReadSeeker) (Metadata, error) {
  m := Metadata{}
  rs.Seek(3, os.SEEK_SET)
  header := make([]byte, 10)
  if _, err := io.ReadFull(rs, header); err != nil {
    return m, err
  }
  h := readHeader(header)
  if h.flags.extendedHeader {

  }
  fmt.Println(h)
  return m, nil
}
