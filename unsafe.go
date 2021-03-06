/* Copyright 2017 Google Inc.
 * https://github.com/cpcallen/flatpack
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

package flatpack

import (
	"reflect"
	"unsafe"
)

// defeat takes a reflect.Value that is addressable but not settable
// (or which can't be used to set another) because it is for/from an
// unexported field, and returns a new reflect.Value that can be set
// (or used to set).
func defeat(v reflect.Value) reflect.Value {
	if !v.CanAddr() {
		panic("Can't defeat protection of unadressable value")
	}
	// Be very careful, if modifying the following statement, not to
	// violate the rules for use of unsafe.Pointer.  In particular,
	// the call of v.UnsafeAddr() and cast to unsafe.Pointer() must be
	// part of a single statement.  It is not safe to temporarily store
	// the intervening uintptr value in a variable.
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}
