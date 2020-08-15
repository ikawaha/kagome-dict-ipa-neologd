package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDh, nil
}

func dictMod0IpaNeologdDictDh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
