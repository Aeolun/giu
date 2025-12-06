//go:build windows

package giu

// applySkipTaskbar is a no-op on Windows as frameless floating windows
// typically don't appear in the taskbar.
func applySkipTaskbar(windowPtr uintptr) {
	// No-op: Windows handles this automatically for frameless floating windows
}
