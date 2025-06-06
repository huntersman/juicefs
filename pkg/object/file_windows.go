/*
 * JuiceFS, Copyright 2020 Juicedata, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package object

import (
	"os"
	"time"
)

func getOwnerGroup(info os.FileInfo) (string, string) {
	return "", ""
}

func lookupUser(name string) int {
	return 0
}

func lookupGroup(name string) int {
	return 0
}

func (d *filestore) Chtimes(key string, mtime time.Time) error {
	p := d.path(key)
	return os.Chtimes(p, time.Time{}, mtime)
}
