package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHk, nil
}

func dictMod0IpaNeologdDictHk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
