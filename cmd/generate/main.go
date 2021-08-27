package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := downloadDB(); err != nil {
		log.Fatal(err)
	}
	if getHash("data/qqwry.dat") == getHash("qqwry.dat") {
		os.Exit(5) // no need to update
	}
	if err := copyFile("data/qqwry.dat", "qqwry.dat"); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove("qqwry.dat"); err != nil {
		log.Fatal(err)
	}
}

func downloadDB() error {
	resp, err := http.Get("https://github.com/out0fmemory/qqwry.dat/raw/master/qqwry_lastest.dat")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var file *os.File
	if file, err = os.Create("qqwry.dat"); err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func getHash(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	h := sha256.New()
	if _, err = io.Copy(h, file); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func copyFile(dstName, srcName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()

	var dst *os.File
	if dst, err = os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
