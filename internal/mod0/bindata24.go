package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAyBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAy, nil
}

func dictMod0IpaNeologdDictAy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ay", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
