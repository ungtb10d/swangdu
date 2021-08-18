// +build darwin netbsd freebsd

package analyze

import (
	"os"
	"syscall"
	"time"
)

const devBSize = 512

func setPlatformSpecificAttrs(file *File, f os.FileInfo) {
	switch stat := f.Sys().(type) {
	case *syscall.Stat_t:
		file.Usage = stat.Blocks * devBSize
		file.Mtime = time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec)

		if stat.Nlink > 1 {
			file.Mli = stat.Ino
		}
	}
}

func setDirPlatformSpecificAttrs(dir *Dir, path string) {
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return
	}

	dir.Mtime = time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec)
}
