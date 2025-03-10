package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf16"
)

var coreui_symbol_massive = []uint16{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
	0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
	0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51,
	0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a,
	0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69,
	0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72,
	0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0xc0,
	0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9,
	0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0, 0xd1, 0xd2,
	0xd3, 0xd4, 0xd5, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc,
	0xdd, 0xde, 0xdf, 0xd6, 0x160, 0x178, 0xe1, 0xe9, 0x152,
	0x2d, 0x2b, 0x2f, 0x3d, 0x21, 0x3f, 0x28, 0x29, 0x5b,
	0x5d, 0x7b, 0x7d, 0x3c, 0x3e, 0x23, 0x24, 0x25, 0x26,
	0x22, 0x60, 0x27, 0x7e, 0x5e, 0x7c, 0x5c, 0x40, 0x3a,
	0x3b, 0x2a, 0x2c, 0x2e, 0x5f, 0x30, 0x31, 0x32, 0x33,
	0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43,
	0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c,
	0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55,
	0x56, 0x57, 0x58, 0x59, 0x5a, 0x61, 0x62, 0x63, 0x64,
	0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d,
	0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76,
	0x77, 0x78, 0x79, 0x7a, 0xc0, 0xc1, 0xc2, 0xc3, 0xc4,
	0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd,
	0xce, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd7,
	0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd6,
	0x160, 0x178, 0xe1, 0xe9, 0x152, 0x2d, 0x2b, 0x2f, 0x3d,
	0x21, 0x3f, 0x28, 0x29, 0x5b, 0x5d, 0x7b, 0x7d, 0x3c,
	0x3e, 0x23, 0x24, 0x25, 0x26, 0x22, 0x60, 0x27, 0x7e,
	0x5e, 0x7c, 0x5c, 0x40, 0x3a, 0x3b, 0x2a, 0x2c, 0x2e,
	0x5f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
	0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50,
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59,
	0x5a, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
	0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71,
	0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a,
	0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8,
	0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0, 0xd1,
	0xd2, 0xd3, 0xd4, 0xd5, 0xd7, 0xd8, 0xd9, 0xda, 0xdb,
	0xdc, 0xdd, 0xde, 0xdf, 0xd6, 0x160, 0x178, 0xe1, 0xe9,
	0x152, 0x2d, 0x2b, 0x2f, 0x3d, 0x21, 0x3f, 0x28, 0x29,
	0x5b, 0x5d, 0x7b, 0x7d, 0x3c, 0x3e, 0x23, 0x24, 0x25,
	0x26, 0x22, 0x60, 0x27, 0x7e, 0x5e, 0x7c, 0x5c, 0x40,
	0x3a, 0x3b, 0x2a, 0x2c, 0x2e, 0x5f, 0x30, 0x31, 0x32,
	0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2e, 0x3a,
	0x25, 0x2b, 0x2d, 0x2f}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func ReadUint32(r io.Reader) uint32 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 4)
	return binary.BigEndian.Uint32(buf.Bytes())
}

func ReadUint16(r io.Reader) uint16 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 2)
	return binary.BigEndian.Uint16(buf.Bytes())
}

func main() {
	args := os.Args
	MCDName := args[1]
	mcd, err := os.Open(MCDName)
	check(err)
	text, err := os.Create(MCDName[:len(MCDName)-4] + ".txt")
	log, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	log.WriteString(MCDName + "\n")
	defer mcd.Close()
	defer text.Close()
	defer log.Close()

	// HeaderRead
	text_packs_offset := ReadUint32(mcd)
	text_packs_count := ReadUint32(mcd)
	block2 := ReadUint32(mcd)
	_ = ReadUint32(mcd) //count2
	_ = ReadUint32(mcd) //symbol_block
	symbol_count := ReadUint32(mcd)
	_ = ReadUint32(mcd) //unk_block1
	_ = ReadUint32(mcd) //unk_block1_count
	_ = ReadUint32(mcd) //unk_block2
	_ = ReadUint32(mcd) //unk_block2_count

	//TextUnpack
	char_massive := make([]uint16, symbol_count)
	mcd.Seek(int64(block2), 0)
	for i := 0; i < int(symbol_count); i++ {
		_ = ReadUint16(mcd) //unk1
		char := ReadUint16(mcd)
		_ = ReadUint16(mcd) //unk2
		char_code := ReadUint16(mcd)
		char_massive[char_code] = char
	}
	mcd.Seek(int64(text_packs_offset), 0)
	for j := 0; j < int(text_packs_count); j++ {
		string_blocks_offset := ReadUint32(mcd)
		blocks_count := ReadUint32(mcd)
		mcd.Seek(int64(string_blocks_offset), 0)
		for k := 0; k < int(blocks_count); k++ {
			strings_desc_offset := ReadUint32(mcd)
			strings_count := ReadUint32(mcd)
			_ = ReadUint32(mcd) //unk1
			_ = ReadUint32(mcd) //unk2
			font_type := ReadUint16(mcd)
			if j == int(text_packs_count)-1 && k == int(blocks_count)-1 {
				log.WriteString(strconv.Itoa(int(font_type)))
			} else {
				log.WriteString(strconv.Itoa(int(font_type)) + ",")
			}
			mcd.Seek(int64(strings_desc_offset), 0)
			for l := 0; l < int(strings_count); l++ {
				string_offset := ReadUint32(mcd)
				_ = ReadUint32(mcd) //unk1
				_ = ReadUint32(mcd) //unk2
				symbols_in_string := ReadUint32(mcd)
				mcd.Seek(int64(string_offset), 0)
				for m := 0; m < int(symbols_in_string); m++ {
					c := ReadUint16(mcd)
					if c == 0x8001 {
						font_type = ReadUint16(mcd)
						binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//switch font_type {
						//case 9:
						//	binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//case 10:
						//	binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//case 1:
						//	binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//case 2:
						//	binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//case 0:
						//	binary.Write(text, binary.LittleEndian, utf16.Encode([]rune(" ")))
						//}
					} else if c == 0x8000 {
						if l == int(strings_count)-1 {
							if symbols_in_string == 1 {
								binary.Write(text, binary.LittleEndian, utf16.Encode([]rune("__dummy\n")))
							} else {
								binary.Write(text, binary.LittleEndian, utf16.Encode([]rune("\n")))
							}
						} else {
							binary.Write(text, binary.LittleEndian, utf16.Encode([]rune("{0A}")))
						}
					} else if c >= 0x4000 && c <= 0x4000+uint16(len(coreui_symbol_massive)) {
						c = c - 0x4000
						binary.Write(text, binary.LittleEndian, coreui_symbol_massive[c])
					} else if c == 0x8004 {
						button_code := ReadUint16(mcd)
						binary.Write(text, binary.LittleEndian, utf16.Encode([]rune("^B"+strconv.Itoa(int(button_code))+"^")))
					} else {
						binary.Write(text, binary.LittleEndian, char_massive[c])
					}
				}
				strings_desc_offset = strings_desc_offset + 24
				mcd.Seek(int64(strings_desc_offset), 0)
			}
			string_blocks_offset = string_blocks_offset + 20
			mcd.Seek(int64(string_blocks_offset), 0)
		}
		text_packs_offset = text_packs_offset + 16
		mcd.Seek(int64(text_packs_offset), 0)
	}
	log.WriteString("\n")
}
