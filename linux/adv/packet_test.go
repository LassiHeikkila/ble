package adv

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestFields(t *testing.T) {
	b, err := hex.DecodeString("0201061AFF4C000215FFFFFFFF0E40706AA38255AAB9EA1A154E2D0055F8000909427573323231303206FF9A08018068")
	if err != nil {
		t.Fatal("failed to decode test data:", err)
	}

	p := Packet{
		b: b,
	}

	d := p.Fields(manufacturerData)
	if len(d) != 2 {
		t.Fatal("wrong number of md fields found:", len(d), "expected", 2)
	}
	if len(d[0]) != (0x1A - 1) {
		t.Fatal("byte slice #1 has wrong length", len(d[0]))
	}
	if len(d[1]) != (0x06 - 1) {
		t.Fatal("byte slice #2 has wrong length", len(d[1]))
	}
	want0 := []byte{
		0x4C, 0x00, 0x02, 0x15, 0xFF, 0xFF, 0xFF, 0xFF, 0x0E, 0x40, 0x70, 0x6A,
		0xA3, 0x82, 0x55, 0xAA, 0xB9, 0xEA, 0x1A, 0x15, 0x4E, 0x2D, 0x00, 0x55,
		0xF8,
	}
	if !bytes.Equal(d[0], want0) {
		t.Fatalf("byte slice #1 not equal:\n%s\n%s\n", hex.EncodeToString(d[0]), hex.EncodeToString(want0))
	}

	want1 := []byte{
		0x9A, 0x08, 0x01, 0x80, 0x68,
	}
	if !bytes.Equal(d[1], want1) {
		t.Fatalf("byte slice #1 not equal:\n%s\n%s\n", hex.EncodeToString(d[1]), hex.EncodeToString(want1))
	}
}
