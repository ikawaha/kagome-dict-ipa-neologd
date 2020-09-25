package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFh, nil
}

func dictMod0IpaNeologdDictFh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
