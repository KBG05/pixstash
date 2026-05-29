// Package storage: implements storage functionalities such as file uploads and downloads
package storage

import (
	"context"
	"io"
)

type BlobInfo struct {
	Path string
	Hash string
	Size int64
}

type BlobStore interface{
	 Put(ctx context.Context, r io.Reader, originalName string)(BlobInfo, error)
}

