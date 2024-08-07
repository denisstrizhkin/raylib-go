//go:build linux && wayland && !x11 && !rgfw && !drm && !sdl && !android
// +build linux,wayland,!x11,!rgfw,!drm,!sdl,!android

package rl

/*
#include "external/glfw/src/wl_init.c"
#include "external/glfw/src/wl_monitor.c"
#include "external/glfw/src/wl_window.c"

#cgo linux CFLAGS: -D_GLFW_WAYLAND
#cgo linux LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl
*/
import "C"
