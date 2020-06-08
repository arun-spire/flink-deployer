package operations

import (
	"context"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/bsm/bfs"
	"github.com/bsm/bfs/bfsfs"
	"github.com/stretchr/testify/assert"
)

/*
 * RetrieveLatestSavepoint
 */
func TestRetrieveLatestSavepointShouldReturnAnErrorIfItCannotReadFromDir(t *testing.T) {
	fs := bfs.NewInMem()
	bfs.Register("inmem", func(_ context.Context, u *url.URL) (bfs.Bucket, error) {
		return fs, nil
	})
	operator := RealOperator{
		Filesystem: fs,
	}

	files, err := operator.retrieveLatestSavepoint("inmem://savepoints")

	assert.Equal(t, "", files)
	assert.EqualError(t, err, "open /savepoints: file does not exist")
}

func TestRetrieveLatestSavepointShouldReturnAnTheNewestFile(t *testing.T) {
	t.Run("Using InMemFS", func(t *testing.T) {
		fs := bfs.NewInMem()
		f1, _ := fs.Create(context.Background(), "savepoint-683b3f-59401d30cfc4/_metadata", nil)
		defer f1.Discard()
		f1.Write([]byte("file a"))
		f1.Commit()
		f2, _ := fs.Create(context.Background(), "savepoint-323b3f-59401d30eoe6/_metadata", nil)
		defer f2.Discard()
		f2.Write([]byte("file b"))
		f2.Commit()
		bfs.Register("inmem", func(_ context.Context, u *url.URL) (bfs.Bucket, error) {
			return fs, nil
		})

		operator := RealOperator{
			Filesystem: fs,
		}

		files, err := operator.retrieveLatestSavepoint("inmem://")

		assert.Equal(t, "inmem://savepoint-323b3f-59401d30eoe6", files)
		assert.Nil(t, err)
	})

	t.Run("Using os FS", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "TestRetrieveLatestSavepointShouldReturnAnTheNewestFile")
		defer os.RemoveAll(dir)
		assert.NoError(t, err)

		fs, err := bfsfs.New(dir, "")
		f1, _ := fs.Create(context.Background(), "savepoint-683b3f-59401d30cfc4/_metadata", nil)
		defer f1.Discard()
		f1.Write([]byte("file a"))
		f1.Commit()
		f2, _ := fs.Create(context.Background(), "savepoint-323b3f-59401d30eoe6/_metadata", nil)
		defer f2.Discard()
		f2.Write([]byte("file b"))
		f2.Commit()

		operator := RealOperator{
			Filesystem: fs,
		}

		files, err := operator.retrieveLatestSavepoint(dir)

		assert.Equal(t, dir+"/savepoint-323b3f-59401d30eoe6", files)
		assert.Nil(t, err)
	})
}

func TestRetrieveLatestSavepointShouldRemoveTheTrailingSlashFromTheSavepointDirectory(t *testing.T) {
	fs := bfs.NewInMem()
	f1, _ := fs.Create(context.Background(), "savepoint-683b3f-59401d30cfc4/_metadata", nil)
	defer f1.Discard()
	f1.Write([]byte("file a"))
	f1.Commit()
	f2, _ := fs.Create(context.Background(), "savepoint-323b3f-59401d30eoe6/_metadata", nil)
	defer f2.Discard()
	f2.Write([]byte("file b"))
	f2.Commit()
	bfs.Register("inmem", func(_ context.Context, u *url.URL) (bfs.Bucket, error) {
		return fs, nil
	})

	operator := RealOperator{
		Filesystem: fs,
	}

	files, err := operator.retrieveLatestSavepoint("inmem://")

	assert.Equal(t, "inmem://savepoint-323b3f-59401d30eoe6", files)
	assert.Nil(t, err)
}

func TestRetrieveLatestSavepointShouldReturnAnErrorWhenDirEmpty(t *testing.T) {
	fs := bfs.NewInMem()
	operator := RealOperator{
		Filesystem: fs,
	}
	bfs.Register("inmem", func(_ context.Context, u *url.URL) (bfs.Bucket, error) {
		return fs, nil
	})

	files, err := operator.retrieveLatestSavepoint("inmem://")

	assert.Equal(t, "", files)
	assert.EqualError(t, err, "No savepoints present in directory: inmem:/")
}
