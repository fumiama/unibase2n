//go:build amd64
// +build amd64

package unibase2n

func cpuid(op uint32) (eax, ebx, ecx, edx uint32)

// True when SSE2 instructions are available.
var canusesse2 = func() bool {
	_, _, c, _ := cpuid(1)
	return c&(1<<26) > 0
}()
