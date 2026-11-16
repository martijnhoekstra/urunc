// Copyright (c) 2023-2026, Nubificus LTD
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

package hypervisors

import (
	"os/exec"

	"github.com/urunc-dev/urunc/pkg/unikontainers/types"
)

const (
	SptVmm    VmmType = "spt"
	SptBinary string  = "solo5-spt"
)

type SPT struct {
	solo5Tender
}

// UsesKVM returns a bool value depending on if the monitor uses KVM
func (s *SPT) UsesKVM() bool {
	return false
}

// Ok checks if the spt binary is available in the system's PATH.
func (s *SPT) Ok() error {
	if _, err := exec.LookPath(SptBinary); err != nil {
		return ErrVMMNotInstalled
	}
	return nil
}

// PreExec performs pre-execution setup. SPT has no special pre-exec requirements.
func (s *SPT) PreExec(_ types.ExecArgs) error {
	return nil
}
