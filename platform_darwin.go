//go:build darwin

package giu

// applySkipTaskbar is a no-op on macOS as frameless floating windows
// automatically don't appear in the Dock.
func applySkipTaskbar(windowPtr uintptr) {
	// No-op: macOS handles this automatically for frameless floating windows
}
