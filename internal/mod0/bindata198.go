package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHq, nil
}

func dictMod0IpaNeologdDictHq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
