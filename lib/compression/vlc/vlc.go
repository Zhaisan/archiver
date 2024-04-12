package vlc

import (
	"log"
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 10010101
	bStr := encodeBin(str)

	// split binary by chunks (8): bits to bytes -> '10010101 10010101 10010101'
	chunks := splitByChunks(bStr, chunksSize)

	// bytes to hex: '20 30 3C'

	return chunks.Bytes()
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).Join()

	// building decoding tree
	// bString (dTree) -> text
	dTree := getEncodingTable().DecodingTree()

	// return decoded text
	return exportText(dTree.Decode(bString))
}

// encodeBin encodes str into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		log.Fatalf("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		's': "0101",
		'r': "01000",
		'h': "0011",
		'l': "001001",
		'd': "00101",
		'c': "000101",
		'u': "00011",
		'm': "000011",
		'f': "000100",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'y': "0000001",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'x': "00000000001",
		'j': "000000001",
		'q': "000000000001",
		'z': "000000000000",
		'!': "001000",
	}
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// exportText is opposite to prepareText, it prepares decoded text to expert
// it changes: ! + <lower case letter> -> to upper case letter
// i.g.: !my name is !ted -> My name is Ted
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false

			continue
		}

		if ch == '!' {
			isCapital = true
			continue
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
