package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDv, nil
}

func dictMod0IpaNeologdDictDv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
