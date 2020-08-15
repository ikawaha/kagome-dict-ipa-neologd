package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDa, nil
}

func dictMod0IpaNeologdDictDa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.da", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
