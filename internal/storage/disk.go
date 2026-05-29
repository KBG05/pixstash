package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

type DiskStore struct{
	Root string
}

func (d *DiskStore) Put(ctx context.Context, r io.Reader, originalName string)(BlobInfo, error){
	// Implementation for storing the file on disk
	hasher:=sha256.New()
	r=io.TeeReader(r, hasher)
	temp, err:=os.CreateTemp(d.Root, "staging-*.tmp")
	if err!=nil{
		return BlobInfo{}, err
	}
	defer temp.Close()
	 _, err=io.Copy(temp, r)
	if err!=nil{
		return BlobInfo{}, err
	}
	hashSum:=hex.EncodeToString(hasher.Sum(nil))
	filepath:=filepath.Join(d.Root, hashSum)
	err=os.Rename(temp.Name(), filepath)
	if err!=nil{
		return BlobInfo{}, err
	}
	info, err:=os.Stat(filepath)
	if err!=nil{
		return BlobInfo{}, err
	}
	return BlobInfo{ Path: filepath, Hash: hashSum, Size: info.Size()}, nil
}
