//go:build !windows

package link

import (
	"errors"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

var haveProgAttach = internal.NewFeatureTest("BPF_PROG_ATTACH", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Type:    ebpf.CGroupSKB,
		License: "MIT",
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
	})
	if err != nil {
		return internal.ErrNotSupported
	}

	// BPF_PROG_ATTACH was introduced at the same time as CGgroupSKB,
	// so being able to load the program is enough to infer that we
	// have the syscall.
	prog.Close()
	return nil
}, "4.10")

var haveProgAttachReplace = internal.NewFeatureTest("BPF_PROG_ATTACH atomic replacement of MULTI progs", func() error {
	if err := haveProgAttach(); err != nil {
		return err
	}

	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Type:       ebpf.CGroupSKB,
		AttachType: ebpf.AttachCGroupInetIngress,
		License:    "MIT",
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
	})

	if err != nil {
		return internal.ErrNotSupported
	}

	defer prog.Close()

	// We know that we have BPF_PROG_ATTACH since we can load CGroupSKB programs.
	// If passing BPF_F_REPLACE gives us EINVAL we know that the feature isn't
	// present.
	attr := sys.ProgAttachAttr{
		// We rely on this being checked after attachFlags.
		TargetFdOrIfindex: ^uint32(0),
		AttachBpfFd:       uint32(prog.FD()),
		AttachType:        uint32(ebpf.AttachCGroupInetIngress),
		AttachFlags:       uint32(flagReplace),
	}

	err = sys.ProgAttach(&attr)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	if errors.Is(err, unix.EBADF) {
		return nil
	}
	return err
}, "5.5")

var haveBPFLink = internal.NewFeatureTest("bpf_link", func() error {
	attr := sys.LinkCreateAttr{
		// This is a hopefully invalid file descriptor, which triggers EBADF.
		TargetFd:   ^uint32(0),
		ProgFd:     ^uint32(0),
		AttachType: sys.AttachType(ebpf.AttachCGroupInetIngress),
	}
	_, err := sys.LinkCreate(&attr)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	if errors.Is(err, unix.EBADF) {
		return nil
	}
	return err
}, "5.7")

var haveProgQuery = internal.NewFeatureTest("BPF_PROG_QUERY", func() error {
	attr := sys.ProgQueryAttr{
		// We rely on this being checked during the syscall.
		// With an otherwise correct payload we expect EBADF here
		// as an indication that the feature is present.
		TargetFdOrIfindex: ^uint32(0),
		AttachType:        sys.AttachType(ebpf.AttachCGroupInetIngress),
	}

	err := sys.ProgQuery(&attr)

	if errors.Is(err, unix.EBADF) {
		return nil
	}
	if err != nil {
		return ErrNotSupported
	}
	return errors.New("syscall succeeded unexpectedly")
}, "4.15")

var haveTCX = internal.NewFeatureTest("tcx", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Type:    ebpf.SchedCLS,
		License: "MIT",
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
	})

	if err != nil {
		return internal.ErrNotSupported
	}

	defer prog.Close()
	attr := sys.LinkCreateTcxAttr{
		// We rely on this being checked during the syscall.
		// With an otherwise correct payload we expect ENODEV here
		// as an indication that the feature is present.
		TargetIfindex: ^uint32(0),
		ProgFd:        uint32(prog.FD()),
		AttachType:    sys.AttachType(ebpf.AttachTCXIngress),
	}

	_, err = sys.LinkCreateTcx(&attr)

	if errors.Is(err, unix.ENODEV) {
		return nil
	}
	if err != nil {
		return ErrNotSupported
	}
	return errors.New("syscall succeeded unexpectedly")
}, "6.6")

var haveNetkit = internal.NewFeatureTest("netkit", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Type:    ebpf.SchedCLS,
		License: "MIT",
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
	})

	if err != nil {
		return internal.ErrNotSupported
	}

	defer prog.Close()
	attr := sys.LinkCreateNetkitAttr{
		// We rely on this being checked during the syscall.
		// With an otherwise correct payload we expect ENODEV here
		// as an indication that the feature is present.
		TargetIfindex: ^uint32(0),
		ProgFd:        uint32(prog.FD()),
		AttachType:    sys.AttachType(ebpf.AttachNetkitPrimary),
	}

	_, err = sys.LinkCreateNetkit(&attr)

	if errors.Is(err, unix.ENODEV) {
		return nil
	}
	if err != nil {
		return ErrNotSupported
	}
	return errors.New("syscall succeeded unexpectedly")
}, "6.7")
