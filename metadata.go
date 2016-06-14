package mm

import (
  "fmt"
)

// Metadata is the construct that holds metadata information about a specific
// audio file
type Metadata map[MetadataField]string

// MetadataField is a type of field that can be found in metadata. Common things
// like Title, or Artist
type MetadataField int8

const (
  TitleField MetadataField = 0
  ArtistField MetadataField = 1
  AlbumField MetadataField = 2
  AlbumArtistField MetadataField = 3
  YearField MetadataField = 4

  // GroupingField describes the music label or producers of an audio file
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

// SetField sets a specific field in Metadata
func (m Metadata) SetField (f MetadataField, v string) {
  m[f] = v
}

// GetField safely retrieves a metadata field
func (m Metadata) GetField (f MetadataField) string {
  v := ""
  if val, ok := m[f]; ok {
    return val
  }
  return v
}

// HasField tests whether the metadata has a field set or not
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
