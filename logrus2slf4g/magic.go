package logrus

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

var defaultMagic = NewMagic(16)

type Magic []byte

func NewMagic(size int) Magic {
	result := make(Magic, size)
	n, err := rand.Reader.Read(result)
	if err != nil {
		panic(fmt.Errorf("cannot aquire magic for logrus: %v", err))
	}
	if n != size {
		panic(fmt.Errorf("cannot aquire magic for logrus: expectd %d bytes; but got: %d", size, n))
	}
	return result
}

func (instance Magic) Equals(o Magic) bool {
	return bytes.Equal(instance, o)
}

func (instance Magic) String() string {
	return hex.EncodeToString(instance)
}
