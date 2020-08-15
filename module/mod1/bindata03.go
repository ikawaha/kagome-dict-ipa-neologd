package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJs, nil
}

func dictMod1IpaNeologdDictJs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.js", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
