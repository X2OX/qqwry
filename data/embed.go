package data

import (
	"embed"
	"encoding/binary"
)

//go:embed qqwry.dat
var content embed.FS

func init() {
	var err error
	if Data, err = content.ReadFile("qqwry.dat"); err != nil {
		panic(err)
	}

	start = int(binary.LittleEndian.Uint32(Data[:4]))
	end = int(binary.LittleEndian.Uint32(Data[4:8]))
}
