package mm

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

type AudioType uint

const (
	FLAC AudioType = iota
	MPEG
)

var Extractors = map[AudioType]*Extractor{
	FLAC: &FlacExtractor{},
	MPEG: &MpegExtractor{},
}

// Parse takes a filename and automatically looks up an extractor to use
func Parse(rs io.ReadSeeker) (*Metadata, error) {
	extractor, err := detectExtractor(rs)
	if err != nil {
		return new(Metadata), err
	}
	return ParseWithExtractor(rs, extractor)
}

// ParseWithExtractor will extract music metadata out of a file with a specific
// extractor
func ParseWithExtractor(rs io.ReadSeeker, ex Extractor) (*Metadata, error) {
	if _, err := rs.Seek(0, os.SEEK_SET); err != nil {
		return nil, errors.Wrap(err, "error occured while seeking data back to 0 offset")
	}
	return ex.Extract(rs)
}

func detectExtractor(rs io.ReadSeeker) (*Extractor, error) {
	var (
		ex     *Extractor
		err    error
		choice string
		ok     bool
	)
	if choice, err = DetectAudioType(rs); err != nil {
		return nil, errors.Wrap(err, "error occured while detecting audio type")
	}
	if ex, ok = Extractors[choice]; !ok {
		return nil, errors.New(err, "detected an audio type that does not have a corresponding extractor")
	}
	return ex, nil
}
