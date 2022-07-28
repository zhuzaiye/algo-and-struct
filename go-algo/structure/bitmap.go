// DATE: 2022/7/17
// DESC: bitmap series and un-series

package structure

import (
	"strconv"
	"strings"
)

var (
	tA = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}
	tB = [8]byte{254, 253, 251, 247, 239, 223, 191, 127}
)

// newSlice creates a new byte slice with length l (in bits).
// The actual size in bits might be up to 7 bits larger because
// they are stored in a byte slice.
func newSlice(l int) []byte {
	remainder := l % 8
	if remainder != 0 {
		remainder = 1
	}
	return make([]byte, l/8+remainder)
}

// Get returns the value of bit i from map m.
// It doesn't check the bounds of the slice.
func get(m []byte, i int) bool {
	return m[i/8]&tA[i%8] != 0
}

// Set sets bit i of map m to value v.
// It doesn't check the bounds of the slice.
func set(m []byte, i int, v bool) {
	index := i / 8
	bit := i % 8
	if v {
		m[index] = m[index] | tA[bit]
	} else {
		m[index] = m[index] & tB[bit]
	}
}

// Len returns the length (in bits) of the provided byteslice.
// It will always be a multipile of 8 bits.
func length(m []byte) int {
	return len(m) * 8
}

// BitMap defines the stored message struct.
type BitMap struct {
	size int
	bits []byte
}

// NewBitMap creates a new bitmap slice.
func NewBitMap(l int) BitMap {
	return BitMap{
		size: l,
		bits: newSlice(l),
	}
}

// Size return bitmap used length.
func (b *BitMap) Size() int {
	return b.size
}

func (b *BitMap) CheckRange(i int) bool {
	if i > b.size || i < 1 {
		return false
	}
	return true
}

func (b *BitMap) Get(i int) bool {
	if i <= 0 || i > b.size {
		return false
	}
	i--
	return get(b.bits, i)
}

func (b *BitMap) Set(i int, v bool) {
	if i <= 0 || i > b.size {
		return
	}
	i--
	set(b.bits, i, v)
}

func (b *BitMap) SetRange(start, end int) {
	for i := start; i <= end; i++ {
		b.Set(i, true)
	}
}

func (b *BitMap) Count() int {
	count := 0
	for i := 0; i < b.Size(); i++ {
		if b.Get(i) {
			count++
		}
	}
	return count
}

// ToString convert the bitmap to string which is separated by space
// The bitmap string can be stored in database
func (b *BitMap) ToString() string {
	var sb strings.Builder
	for i := 0; i < len(b.bits); i++ {
		sb.WriteString(strconv.Itoa(int(b.bits[i])))
		sb.WriteString(" ")
	}
	return sb.String()
}

// FromString convert the specific string to bitmap structure
func (b *BitMap) FromString(input string) {
	inputSlice := strings.Split(input, " ")
	bits := make([]byte, 0, len(inputSlice))
	for i := 0; i < len(inputSlice)-1; i++ {
		bitInt, _ := strconv.Atoi(inputSlice[i])
		bits = append(bits, uint8(bitInt))
	}
	b.size = length(bits)
	b.bits = bits
}
