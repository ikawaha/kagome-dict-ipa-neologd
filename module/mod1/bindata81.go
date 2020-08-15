package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMs, nil
}

func dictMod1IpaNeologdDictMs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ms", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
