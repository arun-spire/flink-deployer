package operations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bsm/bfs"
	_ "github.com/bsm/bfs/bfsfs"
	_ "github.com/bsm/bfs/bfss3"
)

func (o RealOperator) retrieveLatestSavepoint(dir string) (string, error) {
	if !strings.Contains(dir, "://") {
		dir = "file://" + dir
	}
	fs, err := bfs.Connect(context.Background(), dir)
	fmt.Println(fs)
	if err != nil {
		log.Fatalf("Error in connecting to filesystem %s", dir)
		os.Exit(1)
	}

	if strings.HasSuffix(dir, "/") {
		dir = strings.TrimSuffix(dir, "/")
	}

	filesIterator, err := fs.Glob(context.Background(), "*/_metadata")
	if err != nil {
		return "", err
	}
	defer filesIterator.Close()

	var newestFile string
	var newestTime time.Time
	for filesIterator.Next() {
		filePath := filesIterator.Name()
		fmt.Println(filePath)
		currTime := filesIterator.ModTime()
		if currTime.After(newestTime) {
			newestTime = currTime
			newestFile = filePath
		}
	}

	if newestFile == "" {
		return "", errors.New("No savepoints present in directory: " + dir)
	}

	newestFile = strings.TrimPrefix(dir, "file://") + "/" + strings.TrimSuffix(newestFile, "/_metadata")
	return newestFile, nil
}
