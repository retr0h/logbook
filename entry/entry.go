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

// Package entry provides primitives for keeping a log entry.
// There are two essentials types of information that every log needs: Information
// about your operation and information about the station you contact. For your
// operation record the date, frequency, mode and power output; for the contact
// station record their call sign, the time the contact started and ended, their
// signal report, name and location (QTH). When you enter the date and time,
// Universal Coordinated Time (UTC) or Zulu as it is commonly called, is highly
// recommended. Using UTC eliminates confusion over time zones or daylight saving
// time, but you must remember to change the date at 0000Z, which could be anywhere
// from 4 PM to 7 PM local standard time for a North American station. This is an
// advantage of the computerized logging programs. They keep UTC date and time straight
// automatically. Of course, you are free to use local time as long as you indicate
// this clearly in the log. It is unwise to mix UTC and local times and dates together
// in the log; use one or the other.
// http://www.arrl.org/keeping-a-log
package entry

import (
	"encoding/json"
	"time"
)

var (
	// Unmarshaler function to call, which is swapped out in tests.  This may be the wrong
	// way to accomplsih this.
	Unmarshaler = json.Unmarshal
	// Marshaler function to call, which is swapped out in tests.  This may be the wrong
	// way to accomplsih this.
	Marshaler = json.Marshal
)

// Entry is a struct to store a contact.
// - Users (bucket)
//   - KYYZZ: {Name: "Name Last 1", CreatedAt: Time.Now(), ID: 1}
//   - KZZYY: {Name: "Name Last 2", CreatedAt: Time.Now(), ID: 2}
type Entry struct {
	ID        int       `json:"id"`        // ID Auto incrementing from database.
	CreatedAt time.Time `json:"createdAt"` // CreatedAt time entry was created.
	CallSign  string    `json:"callSign"`  // CallSign of contact.
	Name      string    `json:"name"`      // Name of contact.
}

// NewEntry Sets the initial values and returns new Quotes struct.
func NewEntry(callSign string, name string) *Entry {
	return &Entry{
		CreatedAt: newTimeNow(),
		CallSign:  callSign,
		Name:      name,
	}
}

// Unmarshal the data byte slice and returns an Entry struct.
func Unmarshal(data []byte) (*Entry, error) {
	var le *Entry
	err := Unmarshaler(data, &le)
	if err != nil {
		return nil, err
	}
	return le, nil
}

// Marshal the entry struct and returns a data byte slice.
func (le *Entry) Marshal() ([]byte, error) {
	enc, err := Marshaler(le)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func newTimeNow() time.Time {
	return time.Now().UTC()
}
