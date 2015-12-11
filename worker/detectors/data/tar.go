// Copyright 2015 clair authors
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

package data

import (
	"io"
	"strings"

	"github.com/coreos/clair/utils"
	"github.com/coreos/clair/worker/detectors"
)

// TarDataDetector implements DataDetector and detects layer data in 'tar' format
type TarDataDetector struct{}

func init() {
	detectors.RegisterDataDetector("tar", &TarDataDetector{})
}

func (detector *TarDataDetector) Supported(path string, format string) bool {
	switch format {
	case "":
		if strings.HasSuffix(path, ".tar") || strings.HasSuffix(path, ".tar.gz") {
			return true
		}
	case "tar":
		return true
	}
	return false
}

func (detector *TarDataDetector) Detect(layerReader io.ReadCloser, toExtract []string, maxFileSize int64) (map[string][]byte, error) {
	return utils.SelectivelyExtractArchive(layerReader, "./", toExtract, maxFileSize)
}
