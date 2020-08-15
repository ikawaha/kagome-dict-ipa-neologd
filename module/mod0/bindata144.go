package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFo, nil
}

func dictMod0IpaNeologdDictFo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
