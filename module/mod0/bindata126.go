package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEwBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEw, nil
}

func dictMod0IpaNeologdDictEw() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ew", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
