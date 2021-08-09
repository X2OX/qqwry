package main

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	local := getLocalHash()
	remote := getRemoteHash()
	if local == remote {
		log.Println("no need to update")
		return
	}

	b, err := downloadDB()
	if err != nil {
		log.Fatal(err)
	}

	arg := &QQwry{
		Hash:    remote,
		Content: base64.StdEncoding.EncodeToString(b),
		Start:   binary.LittleEndian.Uint32(b[:4]),
		End:     binary.LittleEndian.Uint32(b[4:8]),
	}

	var file *os.File
	if file, err = os.OpenFile("data/data.gen.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
		log.Fatal(err)
	}

	if err = template.Must(template.New("queue").Parse(tpl)).
		Execute(file, arg); err != nil {
		log.Fatal(err)
	}
}

type QQwry struct {
	Hash       string
	Content    string
	Start, End uint32
}

const (
	tpl = `// Code generated, DO NOT EDIT.
// Version: {{.Hash}}

package data

import (
	"encoding/base64"
)

var Data []byte

func init() {
	Data, _ = base64.StdEncoding.DecodeString(__rawData__)
}

const (
	start = {{.Start}}
	end = {{.End}}
	__rawData__ = "{{.Content}}"
)`
)

type githubRepoCommitsHash struct {
	Sha string `json:"sha"`
}

func getRemoteHash() string {
	resp, err := http.Get("https://api.github.com/repos/out0fmemory/qqwry.dat/commits?per_page=1")
	if err != nil {
		return ""
	}
	var arg []githubRepoCommitsHash
	if err = json.NewDecoder(resp.Body).Decode(&arg); err != nil || len(arg) == 0 {
		return ""
	}
	return arg[0].Sha
}
func getLocalHash() string {
	file, err := os.Open("data/data/gen.go")
	if err != nil {
		return ""
	}
	defer file.Close()

	fc := bufio.NewScanner(file)
	for i := 0; fc.Scan(); i++ {
		if i == 1 {
			if h := fc.Text(); len(h) >= 52 {
				return h[13:52]
			}
			return ""
		}
	}
	return ""
}

func downloadDB() ([]byte, error) {
	resp, err := http.Get("https://github.com/out0fmemory/qqwry.dat/raw/master/qqwry_lastest.dat")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
