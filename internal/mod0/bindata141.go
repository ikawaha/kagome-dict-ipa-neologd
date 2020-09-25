package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFl, nil
}

func dictMod0IpaNeologdDictFl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
