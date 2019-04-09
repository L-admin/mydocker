package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func FindCgroupMountpoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo") // 记录当前系统所有挂载文件系统的信息
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	cgroupCompletePath := path.Join(cgroupRoot, cgroupPath)
	if _, err := os.Stat(cgroupCompletePath); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(cgroupCompletePath, 0755); err != nil {
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return cgroupCompletePath, nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}
