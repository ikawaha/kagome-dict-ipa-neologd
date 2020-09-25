package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFi, nil
}

func dictMod0IpaNeologdDictFi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
