package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictInBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIn, nil
}

func dictMod0IpaNeologdDictIn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.in", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
