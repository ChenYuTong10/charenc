package charenc

import (
	"github.com/ChenYuTong10/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

// ToAnsi dispatches to different handler according to the encoding.
// If the encoding has been Ansi, it returns directly. If the encoding is
// out of ANSI, UTF8, BOM UTF8, UTF16 BE/LE, it returns Unexpected Error.
func ToAnsi(stream []byte, encoding string) ([]byte, error) {
	switch strings.ToUpper(encoding) {
	case chardet.BOM_UTF8:
		return BomUTF8ToAnsi(stream)
	case chardet.ANSI:
		return stream, nil
	case chardet.BOM_UTF16_BE:
		return UTF16BEToAnsi(stream)
	case chardet.BOM_UTF16_LE:
		return UTF16LEToAnsi(stream)
	case chardet.UTF8:
		return UTF8ToAnsi(stream)
	default:
		return nil, UnsupportedEncoding(encoding)
	}
}

// BomUTF8ToAnsi transforms stream encoding from BomUTF8 to Ansi.
func BomUTF8ToAnsi(stream []byte) ([]byte, error) {
	stream, _ = BomUTF8ToUTF8(stream)
	return UTF8ToAnsi(stream)
}

// UTF16BEToAnsi transforms stream encoding from UTF16BE to Ansi.
func UTF16BEToAnsi(stream []byte) ([]byte, error) {
	decodeStream, err := UTF16BEToUTF8(stream)
	if err != nil {
		return nil, err
	}
	return UTF8ToAnsi(decodeStream)
}

// UTF16LEToAnsi transforms stream encoding from UTF16LE to Ansi.
func UTF16LEToAnsi(stream []byte) ([]byte, error) {
	decodeStream, err := UTF16LEToUTF8(stream)
	if err != nil {
		return nil, err
	}
	return UTF8ToAnsi(decodeStream)
}

// UTF8ToAnsi transforms stream encoding from UTF8 to Ansi.
func UTF8ToAnsi(stream []byte) ([]byte, error) {
	return simplifiedchinese.GBK.NewEncoder().Bytes(stream)
}
