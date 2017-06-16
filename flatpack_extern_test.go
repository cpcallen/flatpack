/* Copyright 2017 Google Inc.
 * https://github.com/NeilFraser/CodeCity
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package flatpack_test

import (
	"encoding/json"

	"github.com/cpcallen/flatpack"
	"github.com/cpcallen/testutil"
)

func Example() {
	// Some complicated data with shared substructure and cycles:
	type cons struct{ car, cdr interface{} }
	var m = make(map[string]interface{})
	var stuff = &cons{
		car: cons{car: "hello", cdr: 42},
		cdr: &cons{car: m, cdr: nil},
	}
	m["foo"] = stuff
	m["bar"] = stuff
	stuff.cdr.(*cons).cdr = stuff.cdr

	// Create Flatpack and store stuff in it:
	var f = flatpack.New()
	f.Pack("stuff", stuff)
	f.Seal()

	// Convert it to JSON:
	b, e := json.MarshalIndent(f, "", "  ")
	if e != nil {
		panic(e)
	}

	// Register types of stuff we will ask Unpack() to reconstruct
	// (this must include all types it may find in interfaces):
	flatpack.RegisterTypeOf(cons{})
	flatpack.RegisterTypeOf(m)

	// Convert JSON back to Flatpack:
	var f2 *flatpack.Flatpack // N.B.: not using New()
	e = json.Unmarshal(b, &f2)
	if e != nil {
		panic(e)
	}

	// Unpack stuff:
	v, err := f.Unpack("stuff")
	if err != nil {
		panic(err)
	}
	var stuff2 *cons = v.(*cons)

	// Verify stuff2 is an exact (but disjoint) copy of stuff:
	if !testutil.RecEqual(stuff, stuff2, true) {
		panic(testutil.Diff(stuff, stuff2))
	}

	// Output:
}
