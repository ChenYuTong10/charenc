package charenc

import (
	"github.com/ChenYuTong10/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"strings"
)

// ToUTF8 dispatches to different handler according to the encoding.
// If the encoding has been UTF8, it returns directly. If the encoding is
// out of ANSI, UTF8, BOM UTF8, UTF16 BE/LE, it returns Unexpected Error.
func ToUTF8(stream []byte, encoding string) ([]byte, error) {
	switch strings.ToUpper(encoding) {
	case chardet.BOM_UTF8:
		return BomUTF8ToUTF8(stream)
	case chardet.ANSI:
		decodeStream, err := AnsiToUTF8(stream)
		if err != nil {
			return nil, err
		}
		return decodeStream, nil
	case chardet.BOM_UTF16_BE:
		decodeStream, err := UTF16BEToUTF8(stream)
		if err != nil {
			return nil, err
		}
		return decodeStream, nil
	case chardet.BOM_UTF16_LE:
		decodeStream, err := UTF16LEToUTF8(stream)
		if err != nil {
			return nil, err
		}
		return decodeStream, nil
	case chardet.UTF8:
		return stream, nil
	default:
		return nil, UnsupportedEncoding(encoding)
	}
}

// AnsiToUTF8 transforms encoding from Ansi to UTF8.
func AnsiToUTF8(stream []byte) ([]byte, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	return decoder.Bytes(stream)
}

// UTF16BEToUTF8 transforms encoding from UTF16 BE to UTF8.
func UTF16BEToUTF8(stream []byte) ([]byte, error) {
	decoder := unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM).NewDecoder()
	return decoder.Bytes(stream)
}

// UTF16LEToUTF8 transforms encoding from UTF16 LE to UTF8.
func UTF16LEToUTF8(stream []byte) ([]byte, error) {
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewDecoder()
	return decoder.Bytes(stream)
}

// BomUTF8ToUTF8 cut first three bytes BOM prefix of the stream.
func BomUTF8ToUTF8(stream []byte) ([]byte, error) {
	return stream[3:], nil
}
