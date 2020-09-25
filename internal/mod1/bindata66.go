package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMdBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMd, nil
}

func dictMod1IpaNeologdDictMd() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.md", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
