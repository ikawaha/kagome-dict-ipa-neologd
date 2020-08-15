package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFv, nil
}

func dictMod0IpaNeologdDictFv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
