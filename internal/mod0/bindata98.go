package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDu, nil
}

func dictMod0IpaNeologdDictDu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.du", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
