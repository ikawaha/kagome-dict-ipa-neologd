package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDi, nil
}

func dictMod0IpaNeologdDictDi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.di", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
