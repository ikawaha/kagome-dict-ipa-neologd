package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFr, nil
}

func dictMod0IpaNeologdDictFr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
