//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package fd

// GetNumFDs returns the File Descriptors limit.
func GetNumFDs() uint64 { _ = "STUB: not implemented"; return 0 }
