// +build linux

package fs

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/opencontainers/runc/libcontainer/cgroups"
	"github.com/opencontainers/runc/libcontainer/configs"
)

func TestInvalidCgroupPath(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Path: "../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

func TestInvalidAbsoluteCgroupPath(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Path: "/../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidCgroupParent(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "../../../../../../../../../../some/path",
		Name:   "name",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidAbsoluteCgroupParent(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "/../../../../../../../../../../some/path",
		Name:   "name",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidCgroupName(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "parent",
		Name:   "../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}

}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidAbsoluteCgroupName(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "parent",
		Name:   "/../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidCgroupNameAndParent(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "../../../../../../../../../../some/path",
		Name:   "../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

// XXX: Remove me after we get rid of configs.Cgroup.Name and configs.Cgroup.Parent.
func TestInvalidAbsoluteCgroupNameAndParent(t *testing.T) {
	if cgroups.IsCgroup2UnifiedMode() {
		t.Skip("cgroup v2 is not supported")
	}
	root, err := getCgroupRoot()
	if err != nil {
		t.Errorf("couldn't get cgroup root: %v", err)
	}

	config := &configs.Cgroup{
		Parent: "/../../../../../../../../../../some/path",
		Name:   "/../../../../../../../../../../some/path",
	}

	data, err := getCgroupData(config, 0)
	if err != nil {
		t.Errorf("couldn't get cgroup data: %v", err)
	}

	// Make sure the final innerPath doesn't go outside the cgroup mountpoint.
	if strings.HasPrefix(data.innerPath, "..") {
		t.Errorf("SECURITY: cgroup innerPath is outside cgroup mountpoint!")
	}

	// Double-check, using an actual cgroup.
	deviceRoot := filepath.Join(root, "devices")
	devicePath, err := data.path("devices")
	if err != nil {
		t.Errorf("couldn't get cgroup path: %v", err)
	}
	if !strings.HasPrefix(devicePath, deviceRoot) {
		t.Errorf("SECURITY: cgroup path() is outside cgroup mountpoint!")
	}
}

func TestTryDefaultCgroupRoot(t *testing.T) {
	res := tryDefaultCgroupRoot()
	exp := defaultCgroupRoot
	if cgroups.IsCgroup2UnifiedMode() {
		// checking that tryDefaultCgroupRoot does return ""
		// in case /sys/fs/cgroup is not cgroup v1 root dir.
		exp = ""
	}
	if res != exp {
		t.Errorf("tryDefaultCgroupRoot: want %q, got %q", exp, res)
	}
}
