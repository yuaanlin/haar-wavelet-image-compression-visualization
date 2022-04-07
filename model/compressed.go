package model

import (
	"encoding/binary"
	"unsafe"
)

type Compressed struct {
	Width  uint32
	Height uint32
	Level  uint8
	RCount uint32
	GCount uint32
	BCount uint32
	Data   []uint8
	R      []Point
	G      []Point
	B      []Point
}

type Point struct {
	X     uint32
	Y     uint32
	Value uint8
}

func (compressed *Compressed) Bytes() []byte {
	var b []byte
	b = append(b, (*[4]byte)(unsafe.Pointer(&compressed.Width))[:]...)
	b = append(b, (*[4]byte)(unsafe.Pointer(&compressed.Height))[:]...)
	b = append(b, (*[1]byte)(unsafe.Pointer(&compressed.Level))[:]...)
	b = append(b, (*[4]byte)(unsafe.Pointer(&compressed.RCount))[:]...)
	b = append(b, (*[4]byte)(unsafe.Pointer(&compressed.GCount))[:]...)
	b = append(b, (*[4]byte)(unsafe.Pointer(&compressed.BCount))[:]...)
	for _, v := range compressed.Data {
		b = append(b, (*[1]byte)(unsafe.Pointer(&v))[:]...)
	}
	for _, v := range compressed.R {
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.X))[:]...)
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.Y))[:]...)
		b = append(b, (*[1]byte)(unsafe.Pointer(&v.Value))[:]...)
	}
	for _, v := range compressed.G {
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.X))[:]...)
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.Y))[:]...)
		b = append(b, (*[1]byte)(unsafe.Pointer(&v.Value))[:]...)
	}
	for _, v := range compressed.B {
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.X))[:]...)
		b = append(b, (*[4]byte)(unsafe.Pointer(&v.Y))[:]...)
		b = append(b, (*[1]byte)(unsafe.Pointer(&v.Value))[:]...)
	}
	return b
}

func (compressed *Compressed) RGB() [][]RGB {
	result := make([][]RGB, int(compressed.Height))
	for i := 0; i < int(compressed.Height); i++ {
		result[i] = make([]RGB, int(compressed.Width))
		for j := 0; j < int(compressed.Width); j++ {
			result[i][j] = RGB{
				R: 128,
				G: 128,
				B: 128,
			}
		}
	}
	dataWidth := int(compressed.Width) / (int(compressed.Level) + 1)
	dataHeight := int(compressed.Height) / (int(compressed.Level) + 1)
	for i := 0; i < dataHeight; i++ {
		for j := 0; j < dataWidth; j++ {
			result[i][j] = RGB{
				R: int(compressed.Data[3*(i*dataWidth+j)]),
				G: int(compressed.Data[3*(i*dataWidth+j)+1]),
				B: int(compressed.Data[3*(i*dataWidth+j)+2]),
			}
		}
	}
	for i := 0; i < int(compressed.RCount); i++ {
		p := compressed.R[i]
		result[int(p.Y)][int(p.X)].R = int(p.Value)
	}
	for i := 0; i < int(compressed.GCount); i++ {
		p := compressed.G[i]
		result[int(p.Y)][int(p.X)].G = int(p.Value)
	}
	for i := 0; i < int(compressed.BCount); i++ {
		p := compressed.B[i]
		result[int(p.Y)][int(p.X)].B = int(p.Value)
	}
	return result
}

func (compressed *Compressed) FromBytes(b []byte) {
	c := compressed
	c.Width = binary.LittleEndian.Uint32(b[0:4])
	c.Height = binary.LittleEndian.Uint32(b[4:8])
	c.Level = b[8]
	c.RCount = binary.LittleEndian.Uint32(b[9:13])
	c.GCount = binary.LittleEndian.Uint32(b[13:17])
	c.BCount = binary.LittleEndian.Uint32(b[17:21])
	dataHeight := int(c.Height) / (int(c.Level) + 1)
	dataWidth := int(c.Width) / (int(c.Level) + 1)
	for i := 0; i < 3*dataWidth*dataHeight; i++ {
		c.Data = append(c.Data, b[21+i])
	}
	s := 21 + 3*dataWidth*dataHeight
	for i := 0; i < int(c.RCount); i++ {
		x := binary.LittleEndian.Uint32(b[s+9*i : s+9*i+4])
		y := binary.LittleEndian.Uint32(b[s+9*i+4 : s+9*i+8])
		c.R = append(c.R, Point{
			X:     x,
			Y:     y,
			Value: b[s+9*i+8],
		})
	}
	s = 21 + 3*dataWidth*dataHeight + 9*(int(c.RCount))
	for i := 0; i < int(c.GCount); i++ {
		c.G = append(c.G, Point{
			X:     binary.LittleEndian.Uint32(b[s+9*i : s+9*i+4]),
			Y:     binary.LittleEndian.Uint32(b[s+9*i+4 : s+9*i+8]),
			Value: b[s+9*i+8],
		})
	}
	s = 21 + 3*dataWidth*dataHeight + 9*(int(c.RCount)) + 9*(int(c.GCount))
	for i := 0; i < int(c.BCount); i++ {
		c.B = append(c.B, Point{
			X:     binary.LittleEndian.Uint32(b[s+9*i : s+9*i+4]),
			Y:     binary.LittleEndian.Uint32(b[s+9*i+4 : s+9*i+8]),
			Value: b[s+9*i+8],
		})
	}
}

func (compressed *Compressed) FromRGB(data [][]RGB, level int) {
	compressed.Width = uint32(len(data[0]))
	compressed.Height = uint32(len(data))
	compressed.Level = uint8(level)
	dataWidth := len(data[0]) / (level + 1)
	dataHeight := len(data) / (level + 1)
	for i := 0; i < dataHeight; i++ {
		for j := 0; j < dataWidth; j++ {
			compressed.Data = append(compressed.Data, uint8(data[i][j].R))
			compressed.Data = append(compressed.Data, uint8(data[i][j].G))
			compressed.Data = append(compressed.Data, uint8(data[i][j].B))
		}
	}
	for i, row := range data {
		for j, col := range row {
			if i < dataHeight && j < dataWidth {
				continue
			}
			if col.R != 128 {
				compressed.RCount++
				compressed.R = append(compressed.R, Point{
					X:     uint32(j),
					Y:     uint32(i),
					Value: uint8(col.R),
				})
			}
			if col.G != 128 {
				compressed.GCount++
				compressed.G = append(compressed.G, Point{
					X:     uint32(j),
					Y:     uint32(i),
					Value: uint8(col.G),
				})
			}
			if col.B != 128 {
				compressed.BCount++
				compressed.B = append(compressed.B, Point{
					X:     uint32(j),
					Y:     uint32(i),
					Value: uint8(col.B),
				})
			}
		}
	}
}
