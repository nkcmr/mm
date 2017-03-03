package mm

import (
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
)

var ErrUnsupportedAudioType = errors.New("cannot detect the type of audio in data")

type audioFileSignature struct {
	offset  int64
	markers [][]byte
}

var signatures = map[AudioType]audioFileSignature{
	FLAC: audioFileSignature{
		offset: 0,
		markers: [][]byte{
			[]byte("fLaC"),
		},
	},
	MPEG: audioFileSignature{
		offset: 0,
		markers: [][]byte{
			[]byte("ID32"),
			[]byte("ID33"),
			[]byte("ID34"),
		},
	},
}

// DetectAudioType takes a readSeeker and will detect the type of audio based on
// specific bytes that are present in the first few bytes of a file
func DetectAudioType(rs io.ReadSeeker) (AudioType, error) {
	for kind, sig := range signatures {
		for _, marker := range sig.markers {
			if _, err := rs.Seek(sig.offset, os.SEEK_SET); err != nil {
				return nil, errors.Wrap(err, "error occured while seeking over data")
			}
			data := make([]byte, len(marker))
			if _, err := io.ReadFull(rs, data); err != nil {
				return nil, errors.Wrap(err, "error occured while reading data")
			}
			if bytes.Equal(marker, data) {
				return kind, nil
			}
		}
	}
	return nil, ErrUnsupportedAudioType
}
