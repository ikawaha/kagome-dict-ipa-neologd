package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIu, nil
}

func dictMod0IpaNeologdDictIu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.iu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
