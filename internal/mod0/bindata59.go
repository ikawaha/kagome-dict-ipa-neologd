package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictChBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCh, nil
}

func dictMod0IpaNeologdDictCh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictChBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ch", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
