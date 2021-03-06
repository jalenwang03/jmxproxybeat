package nfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_msg = []byte{
	0x80, 0x00, 0x00, 0xe0,
	0xb5, 0x49, 0x21, 0xab,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	0x00, 0x00, 0x00, 0x04,

	0x00, 0x00, 0x00, 0x0b,
	0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x00,
}

func TestXdrDecoding(t *testing.T) {
	xdr := Xdr{data: test_msg, offset: 0}

	assert.Equal(t, uint32(0x800000e0), uint32(xdr.getUInt()))
	assert.Equal(t, uint32(0xb54921ab), uint32(xdr.getUInt()))
	assert.Equal(t, uint64(2), uint64(xdr.getUHyper()))
	assert.Equal(t, uint32(4), uint32(xdr.getUInt()))
	assert.Equal(t, "test string", xdr.getString())
}
