package host

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

// BackupFile creates a copy of your hosts file to a new location with the date as extension
func BackupFile(src, dstPath string) (string, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	bkpFilename := getBackupFilename(srcFile.Name(), dstPath, time.Now())
	dst, err := os.Create(bkpFilename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, srcFile)
	return bkpFilename, err
}

func getBackupFilename(srcFilename, dstPath string, t time.Time) string {
	bkpFilename := fmt.Sprintf("%s.%s", srcFilename, t.UTC().Format("20060102"))
	bkpFilename = path.Join(dstPath, path.Base(bkpFilename))

	return bkpFilename
}
