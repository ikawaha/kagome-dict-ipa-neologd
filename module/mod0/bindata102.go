package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDyBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDy, nil
}

func dictMod0IpaNeologdDictDy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
