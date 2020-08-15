package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDx, nil
}

func dictMod0IpaNeologdDictDx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
