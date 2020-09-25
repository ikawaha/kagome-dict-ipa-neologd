package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHu, nil
}

func dictMod0IpaNeologdDictHu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
