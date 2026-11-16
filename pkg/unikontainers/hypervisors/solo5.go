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
	"github.com/urunc-dev/urunc/pkg/unikontainers/types"
	"golang.org/x/sys/unix"
)

// solo5ArgsSeparator terminates the tender options in a solo5 cli
const solo5ArgsSeparator = "--"

// solo5Tender holds what the solo5 tenders (hvt and spt) share. HVT and SPT
// embed it and add what differs: KVM use, the binary to look up, and PreExec.
type solo5Tender struct {
	binaryPath string
	binary     string
}

func (t *solo5Tender) Signal(pid int, signal unix.Signal) error {
	return signalProcess(pid, signal)
}

// Stop kills the tender process
func (t *solo5Tender) Stop(pid int) error {
	return killProcess(pid)
}

// UsesTUNTAP reports that the guest gets its network through a tap device.
func (t *solo5Tender) UsesTUNTAP() bool {
	return true
}

// SupportsSharedfs returns a bool value depending on the monitor support for shared-fs
func (t *solo5Tender) SupportsSharedfs(_ string) bool {
	return false
}

// SupportsControlSocket reports that the solo5 tenders expose no control socket.
func (t *solo5Tender) SupportsControlSocket() bool {
	return false
}

// Path returns the path to the tender binary.
func (t *solo5Tender) Path() string {
	return t.binaryPath
}

// BuildExecCmd builds the command line of a solo5 tender. Both tenders share
// the cli grammar "[ OPTIONS ] [ -- ] KERNEL [ ARGS ]".
func (t *solo5Tender) BuildExecCmd(args types.ExecArgs, ukernel types.Unikernel) ([]string, error) {
	exArgs := []string{t.binaryPath, "--mem=" + BytesToStringMB(args.MemSizeB)}
	if args.Net.TapDev != "" {
		netCli := ukernel.MonitorNetCli(args.Net.TapDev, args.Net.MAC)
		exArgs = append(exArgs, netCli...)
	}
	bArgs := ukernel.MonitorBlockCli()
	for _, blockArg := range bArgs {
		if blockArg.Path != "" {
			exArgs = append(exArgs, "--block:"+blockArg.ID+"="+blockArg.Path)
		}
	}
	extraMonArgs := ukernel.MonitorCli()
	exArgs = append(exArgs, extraMonArgs.OtherArgs...)
	exArgs = append(exArgs, solo5ArgsSeparator, args.UnikernelPath)
	if args.Command != "" {
		exArgs = append(exArgs, args.Command)
	}
	return exArgs, nil
}
