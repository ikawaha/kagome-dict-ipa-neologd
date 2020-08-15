package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDwBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDw, nil
}

func dictMod0IpaNeologdDictDw() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
