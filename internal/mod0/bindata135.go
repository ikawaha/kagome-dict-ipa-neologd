package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFf, nil
}

func dictMod0IpaNeologdDictFf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ff", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
