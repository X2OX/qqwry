package data

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	IndexLen      = 7
	RedirectMode1 = 0x01
	RedirectMode2 = 0x02
)

func New() *Cz { return &Cz{} }

type Cz struct {
	ip      uint32
	Error   error  `json:"-"`
	IP      net.IP `json:"ip"`
	Country string `json:"country"`
	Area    string `json:"area"`
}

func (c *Cz) Find(ip string) *Cz {
	fmt.Println(ip)
	return c.FindIP(net.ParseIP(ip))
}

func (c *Cz) FindIP(ip net.IP) *Cz {
	if c.Error = c.parseIP(ip); c.Error != nil {
		return c
	}
	return c.find()
}

func (c *Cz) parseIP(ip net.IP) error {
	c.IP = ip
	if arr := ip.To4(); len(arr) == net.IPv4len {
		c.ip = binary.BigEndian.Uint32(arr)
		return nil
	}
	fmt.Println(ip)
	return errors.New("ip wrong")
}

func (c *Cz) find() *Cz {
	var (
		country, area []byte
		offset        = c.searchRecord(c.ip)
	)

	if offset <= 0 {
		c.Error = errors.New("IP 未找到归属地")
		return c
	}

	switch Data[offset+4] {
	case RedirectMode1:
		countryOffset := readUint32FromByte3(offset + 5)

		switch Data[countryOffset] {
		case RedirectMode2:
			rc := readUint32FromByte3(countryOffset + 1)
			country = c.readString(rc)
			countryOffset += 4
			area = c.readArea(countryOffset)
		default:
			country = c.readString(countryOffset)
			countryOffset += uint32(len(country) + 1)
			area = c.readArea(countryOffset)
		}
	case RedirectMode2:
		countryOffset := readUint32FromByte3(offset + 5)
		country = c.readString(countryOffset)
		area = c.readArea(offset + 8)
	}

	enc := simplifiedchinese.GBK.NewDecoder()
	c.Country, _ = enc.String(string(country))
	c.Area, _ = enc.String(string(area))

	if c.Area == " CZ88.NET" {
		c.Area = ""
	}
	return c
}

func (c *Cz) readString(offset uint32) []byte {
	i := 0
	for ; Data[int(offset)+i] != 0; i++ {
	}
	return Data[offset : int(offset)+i]
}

func (c *Cz) readArea(offset uint32) []byte {
	mode := Data[offset : offset+1][0]
	if mode == RedirectMode1 || mode == RedirectMode2 {
		areaOffset := readUint32FromByte3(offset + 1)
		if areaOffset == 0 {
			return []byte("")
		}
		return c.readString(areaOffset)
	}
	return c.readString(offset)
}

func (c *Cz) searchRecord(ip uint32) uint32 {
	_start, _end := start, end

	for {
		mid := _start + (((_end-_start)/IndexLen)>>1)*IndexLen
		buf := Data[mid : mid+IndexLen]
		_ip := binary.LittleEndian.Uint32(buf[:4])

		if _end-_start == IndexLen {
			offset := byte3ToUInt32(buf[4:7])
			buf = Data[mid+IndexLen : mid+IndexLen+IndexLen]
			if ip < binary.LittleEndian.Uint32(buf[:4]) {
				return offset
			}
			return 0
		}

		if _ip > ip {
			_end = mid
		} else if _ip < ip {
			_start = mid
		} else if _ip == ip {
			return byte3ToUInt32(buf[4:7])
		}
	}
}
