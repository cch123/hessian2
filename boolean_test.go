// Copyright (c) 2016 ~ 2019, Alex Stocks.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hessian

import (
	"testing"
)

func TestEncBool(t *testing.T) {
	var (
		e    *Encoder
		want []byte
	)

	e = NewEncoder()
	e.Encode(true)
	if e.Buffer()[0] != 'T' {
		t.Fail()
	}
	want = []byte{0x54}
	assertEqual(want, e.Buffer(), t)

	e = NewEncoder()
	e.Encode(false)
	if e.Buffer()[0] != 'F' {
		t.Fail()
	}
	want = []byte{0x46}
	assertEqual(want, e.Buffer(), t)
}

func testBooleanFramework(t *testing.T, method string, expected bool) {
	r, e := decodeResponse(method)
	if e != nil {
		t.Errorf("%s: decode fail with error %v", method, e)
		return
	}

	v, ok := r.(bool)
	if !ok {
		t.Errorf("%s: %v is not bool", method, r)
		return
	}

	if ok && v != expected {
		t.Errorf("%s: got %v, wanted %v", method, v, expected)
	}
}

func TestBoolean(t *testing.T) {
	testBooleanFramework(t, "replyFalse", false)
	testBooleanFramework(t, "replyTrue", true)
}
