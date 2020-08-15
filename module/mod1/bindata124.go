package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOjBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOj, nil
}

func dictMod1IpaNeologdDictOj() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
