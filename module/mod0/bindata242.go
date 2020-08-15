package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJi, nil
}

func dictMod0IpaNeologdDictJi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ji", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
