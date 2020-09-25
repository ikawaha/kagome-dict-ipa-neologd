package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHb, nil
}

func dictMod0IpaNeologdDictHb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
