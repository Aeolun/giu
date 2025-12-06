//go:build linux

package giu

/*
#cgo pkg-config: x11
#include <X11/Xlib.h>
#include <X11/Xatom.h>
#include <stdlib.h>

// Forward declare GLFW function
typedef struct GLFWwindow GLFWwindow;
extern unsigned long glfwGetX11Window(GLFWwindow* window);

void setSkipTaskbar(GLFWwindow* glfwWindow) {
	Display* display = XOpenDisplay(NULL);
	if (!display) {
		return;
	}

	Window x11Window = glfwGetX11Window(glfwWindow);
	if (!x11Window) {
		XCloseDisplay(display);
		return;
	}

	// Set window type to notification
	Atom wmWindowType = XInternAtom(display, "_NET_WM_WINDOW_TYPE", False);
	Atom wmWindowTypeNotification = XInternAtom(display, "_NET_WM_WINDOW_TYPE_NOTIFICATION", False);
	XChangeProperty(display, x11Window, wmWindowType, XA_ATOM, 32, PropModeReplace,
		(unsigned char*)&wmWindowTypeNotification, 1);

	// Also set skip taskbar and skip pager states
	Atom wmState = XInternAtom(display, "_NET_WM_STATE", False);
	Atom wmStateSkipTaskbar = XInternAtom(display, "_NET_WM_STATE_SKIP_TASKBAR", False);
	Atom wmStateSkipPager = XInternAtom(display, "_NET_WM_STATE_SKIP_PAGER", False);
	Atom wmStateAbove = XInternAtom(display, "_NET_WM_STATE_ABOVE", False);

	Atom states[3] = {wmStateSkipTaskbar, wmStateSkipPager, wmStateAbove};
	XChangeProperty(display, x11Window, wmState, XA_ATOM, 32, PropModeReplace,
		(unsigned char*)states, 3);

	XFlush(display);
	XCloseDisplay(display);
}
*/
import "C"
import (
	"unsafe"
)

func applySkipTaskbar(windowPtr uintptr) {
	if windowPtr == 0 {
		return
	}

	C.setSkipTaskbar((*C.GLFWwindow)(unsafe.Pointer(windowPtr)))
}
