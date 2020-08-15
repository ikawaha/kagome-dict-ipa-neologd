package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIo, nil
}

func dictMod0IpaNeologdDictIo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.io", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
