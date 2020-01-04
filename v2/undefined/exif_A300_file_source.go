package exifundefined

import (
	"encoding/binary"

	"github.com/dsoprea/go-logging"

	"github.com/dsoprea/go-exif/v2/common"
)

type TagExifA300FileSource uint32

const (
	TagUndefinedType_A300_SceneType_Others                   TagExifA300FileSource = 0
	TagUndefinedType_A300_SceneType_ScannerOfTransparentType TagExifA300FileSource = 1
	TagUndefinedType_A300_SceneType_ScannerOfReflexType      TagExifA300FileSource = 2
	TagUndefinedType_A300_SceneType_Dsc                      TagExifA300FileSource = 3
)

type CodecExifA300FileSource struct {
}

func (CodecExifA300FileSource) Encode(value interface{}, byteOrder binary.ByteOrder) (encoded []byte, unitCount uint32, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	st, ok := value.(TagExifA300FileSource)
	if ok == false {
		log.Panicf("can only encode a TagExifA300FileSource")
	}

	ve := exifcommon.NewValueEncoder(byteOrder)

	ed, err := ve.Encode([]uint32{uint32(st)})
	log.PanicIf(err)

	// TODO(dustin): Confirm this size against the specification. It's non-specific about what type it is, but it looks to be no more than a single integer scalar. So, we're assuming it's a LONG.

	return ed.Encoded, uint32(int(ed.UnitCount)), nil
}

func (CodecExifA300FileSource) Decode(valueContext *exifcommon.ValueContext) (value interface{}, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	valueContext.SetUndefinedValueType(exifcommon.TypeLong)

	valueLongs, err := valueContext.ReadLongs()
	log.PanicIf(err)

	return TagExifA300FileSource(valueLongs[0]), nil
}

func init() {
	registerEncoder(
		TagExifA300FileSource(0),
		CodecExifA300FileSource{})

	registerDecoder(
		exifcommon.IfdPathStandardExif,
		0xa300,
		CodecExifA300FileSource{})
}