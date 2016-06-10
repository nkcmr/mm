package mm

import (
  "io"
  "os"
  "encoding/binary"
  "strings"
  "fmt"
)

type blockType int8

const (
  streaminfo blockType = 0
  padding blockType = 1
  application blockType = 2
  seektable blockType = 3
  vorbisComment blockType = 4
  cuesheet blockType = 5
  picture blockType = 6
  invalidBlock blockType = 127
)

var blockTypeName = map[blockType]string{
  streaminfo: "streaminfo",
  padding: "padding",
  application: "application",
  seektable: "seektable",
  vorbisComment: "vorbis_comment",
  cuesheet: "cuesheet",
  picture: "picture",
  invalidBlock: "invalid",
}

var vorbisFieldMap = map[string]MetadataField{
  "TITLE": TitleField,
  "ARTIST": ArtistField,
  "ALBUM": AlbumField,
  "ALBUMARTIST": AlbumArtistField,
  "DATE": YearField,
  "ORGANIZATION": GroupingField,
  "TRACKNUMBER": TrackNumberField,
  "TRACKTOTAL": TrackTotalField,
  "DISCNUMBER": DiscNumberField,
  "DISCTOTAL": DiscTotalField,
}

func (k blockType) String() string {
  if name, ok := blockTypeName[k]; ok {
    return name
  }
  return "reserved"
}

type FlacExtractor struct {}

func (ex FlacExtractor) Extract(rs io.ReadSeeker) (Metadata, error) {
  m := Metadata{}
  rs.Seek(4, os.SEEK_SET)
  for {
    last, kind, data, err := readMetaDataBlock(rs)
    if err != nil {
      return m, err
    }
    if kind == vorbisComment {
      for field, value := range readVorbisComment(data) {
        fmt.Println(field, value)
        if coreField, ok := vorbisFieldMap[field]; ok {
          m.SetField(coreField, value)
        }
      }
    }
    if last {
      break
    }
  }
  return m, nil
}

func readMetaDataBlock(r io.Reader) (last bool, kind blockType, data []byte, err error) {
  header := make([]byte, 4)
  if _, err := io.ReadFull(r, header); err != nil {
    return last, kind, data, err
  }
  n := binary.BigEndian.Uint32(header)
  last = (n >> 31) == 1
  kind = blockType((n >> 24) ^ (n >> 31 << 7))
  data = make([]byte, int32(n << 8 >> 8))
  if _, err := io.ReadFull(r, data); err != nil {
    return last, kind, data, err
  }
  return last, kind, data, nil
}

func readVorbisComment(blockData []byte) map[string]string {
  info := map[string]string{}

  // read vendor string
  vendorLength := binary.LittleEndian.Uint32(blockData[:4])
  offset := vendorLength + 4
  info["vendor"] = string(blockData[4:offset])

  // begin reading user comment list
  numFields := binary.LittleEndian.Uint32(blockData[offset:(offset + 4)])
  offset += 4
  for i := 0; i < int(numFields); i++ {
    fieldLength := binary.LittleEndian.Uint32(blockData[offset:(offset + 4)])
    offset += 4
    fieldValue := blockData[offset:(offset + fieldLength)]
    offset += fieldLength
    sepIdx := strings.Index(string(fieldValue), "=")
    if sepIdx == -1 {
      continue
    }
    info[string(fieldValue[0:sepIdx])] = string(fieldValue[(sepIdx + 1):])
  }
  return info
}
