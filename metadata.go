package mm

import (
  "fmt"
)

type Metadata map[MetadataField]string

type MetadataField int8

const (
  TitleField MetadataField = 0
  ArtistField MetadataField = 1
  AlbumField MetadataField = 2
  AlbumArtistField MetadataField = 3
  YearField MetadataField = 4
  GroupingField MetadataField = 5

  TrackNumberField MetadataField = 6
  TrackTotalField MetadataField = 7
  DiscNumberField MetadataField = 8
  DiscTotalField MetadataField = 9
)

var fieldNameMap = map[MetadataField]string{
  TitleField: "title",
  ArtistField: "artist",
  AlbumField: "album",
  AlbumArtistField: "album_artist",
  YearField: "year",
  GroupingField: "grouping",

  TrackNumberField: "track_number",
  TrackTotalField: "track_total",
  DiscNumberField: "disc_number",
  DiscTotalField: "disc_total",
}

func (m Metadata) SetField (f MetadataField, v string) {
  m[f] = v
}

func (m Metadata) GetField (f MetadataField) string {
  v := ""
  if val, ok := m[f]; ok {
    return val
  }
  return v
}

func (m Metadata) HasField (f MetadataField) bool {
  _, ok := m[f]
  return ok
}

func (m Metadata) String () string {
  out := "{\n"
  for k, v := range m {
    out += fmt.Sprintf("  %s: %s\n", fieldNameMap[k], v);
  }
  out += "}\n";
  return out
}
