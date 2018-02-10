// Copyright (c) 2018 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package entry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/retr0h/logbook/entry"
	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	callSign := "KYYZZ"
	name := "Test User"
	ne := entry.NewEntry(callSign, name)

	assert.IsType(t, int(0), ne.ID)
	assert.IsType(t, time.Time{}, ne.CreatedAt)
	assert.Equal(t, "KYYZZ", ne.CallSign)
	assert.Equal(t, "Test User", ne.Name)
}

func TestUnmarshal(t *testing.T) {
	callSign := "KYYZZ"
	name := "Test User"
	ne := entry.NewEntry(callSign, name)
	data, _ := ne.Marshal()

	e, _ := entry.Unmarshal(data)
	assert.Equal(t, "KYYZZ", e.CallSign)
	assert.Equal(t, "Test User", e.Name)
}

func TestUnmarshalReturnsErrorsWhenJsonUnmarshalFails(t *testing.T) {
	data := ``
	originalUnmarshaler := entry.Unmarshaler
	entry.Unmarshaler = func([]byte, interface{}) error {
		return errors.New("Failed to Unmarshal")
	}
	defer func() { entry.Unmarshaler = originalUnmarshaler }()

	_, err := entry.Unmarshal([]byte(data))

	assert.Error(t, err)
}

func TestMarshal(t *testing.T) {
	callSign := "KYYZZ"
	name := "Test User"
	ne := entry.NewEntry(callSign, name)
	data, _ := ne.Marshal()

	assert.IsType(t, []byte{}, data)
}

func TestMarshalReturnsErrorsWhenJsonMarshalFails(t *testing.T) {
	callSign := "KYYZZ"
	name := "Test User"
	ne := entry.NewEntry(callSign, name)
	originalMarshaler := entry.Marshaler
	entry.Marshaler = func(interface{}) ([]byte, error) {
		err := errors.New("Failed to Marshal")

		return []byte{}, err
	}
	defer func() { entry.Marshaler = originalMarshaler }()

	_, err := ne.Marshal()

	assert.Error(t, err)
}
