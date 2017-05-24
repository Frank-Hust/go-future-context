//Copyright 2016 Capital One Services, LLC
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and limitations under the License.
// SPDX-Copyright: Copyright (c) Capital One Services, LLC
// SPDX-License-Identifier: Apache-2.0

package future_test

import (
	"fmt"
	"time"

	future "github.com/capitalone/go-future-context"
)

func Example_getUntil() {
	inVal := 200

	ThingThatTakesALongTimeToCalculate := func(inVal int) (int, error) {
		//this does something but it's not that important
		time.Sleep(5 * time.Second)
		return inVal * 2, nil
	}

	f := future.New(func() (interface{}, error) {
		return ThingThatTakesALongTimeToCalculate(inVal)
	})

	// times out
	result, timeout, err := f.GetUntil(2 * time.Second)
	fmt.Println(result, timeout, err)

	// completes before 10s passes
	result, timeout, err = f.GetUntil(10 * time.Second)
	fmt.Println(result, timeout, err)

	//  results are instantaneous
	result, timeout, err = f.GetUntil(10 * time.Second)
	fmt.Println(result, timeout, err)

	// Output:
	// <nil> true <nil>
	// 400 false <nil>
	// 400 false <nil>

}
