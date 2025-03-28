package common

import (
	"encoding/binary"
)

const (
	EntryTypeTopicData                      = 0
	EntryTypeOffsetSnapshot                 = 1
	EntryTypeOffsetTime                     = 2
	EntryTypeCompactedTopicLastOffsetForKey = 3
)

func AppendValueMetadata(buff []byte, meta ...int64) []byte {
	start := len(buff)
	numValues := len(meta)
	for i := 0; i < numValues; i++ {
		buff = binary.AppendVarint(buff, meta[i])
	}
	size := len(buff) - start
	if size > 255 {
		panic("too many values to append to value metadata")
	}
	buff = append(buff, byte(size))
	return buff
}

func RemoveValueMetadata(buff []byte) []byte {
	lb := len(buff)
	size := int(buff[lb-1])
	startPos := lb - size - 1
	return buff[:startPos]
}

func ReadValueMetadata(buff []byte) []int64 {
	lb := len(buff)
	size := int(buff[lb-1])
	startPos := lb - size - 1
	var values []int64
	for startPos < lb-1 {
		val, read := binary.Varint(buff[startPos:])
		values = append(values, val)
		startPos += read
	}
	return values
}

func ReadAndRemoveValueMetadata(buff []byte) ([]int64, []byte) {
	lb := len(buff)
	size := int(buff[lb-1])
	startPos := lb - size - 1
	pos := startPos
	var values []int64
	for pos < lb-1 {
		val, read := binary.Varint(buff[pos:])
		values = append(values, val)
		pos += read
	}
	return values, buff[:startPos]
}
