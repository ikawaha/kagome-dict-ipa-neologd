package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEm, nil
}

func dictMod0IpaNeologdDictEm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.em", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
