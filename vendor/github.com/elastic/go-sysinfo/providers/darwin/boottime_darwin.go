// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build (amd64 && cgo) || (arm64 && cgo)
// +build amd64,cgo arm64,cgo

package darwin

import (
	"syscall"
	"time"

	"github.com/pkg/errors"
)

const kernBoottimeMIB = "kern.boottime"

func BootTime() (time.Time, error) {
	var tv syscall.Timeval
	if err := sysctlByName(kernBoottimeMIB, &tv); err != nil {
		return time.Time{}, errors.Wrap(err, "failed to get host uptime")
	}

	bootTime := time.Unix(int64(tv.Sec), int64(tv.Usec)*int64(time.Microsecond))
	return bootTime, nil
}
