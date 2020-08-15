package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDz, nil
}

func dictMod0IpaNeologdDictDz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
