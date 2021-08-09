package data

func readUint32FromByte3(offset uint32) uint32 {
	return byte3ToUInt32(Data[offset : offset+3])
}

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
