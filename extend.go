package csv

var (
	readQuoteChar     byte = '"'
	runeReadQuoteChar rune = '"'
)

func SetReadQuoteChar(c byte) {
	readQuoteChar = c
	runeReadQuoteChar = rune(c)
}
