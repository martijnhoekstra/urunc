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
	HvtVmm    VmmType = "hvt"
	HvtBinary string  = "solo5-hvt"
)

type HVT struct {
	solo5Tender
}

// UsesKVM returns a bool value depending on if the monitor uses KVM
func (h *HVT) UsesKVM() bool {
	return true
}

// Ok checks if the hvt binary is available in the system's PATH.
func (h *HVT) Ok() error {
	if _, err := exec.LookPath(HvtBinary); err != nil {
		return ErrVMMNotInstalled
	}
	return nil
}

// PreExec performs pre-execution setup for HVT.
// Solo5-hvt applies its own seccomp filter, so there is nothing to do.
func (h *HVT) PreExec(args types.ExecArgs) error {
	return nil
}
