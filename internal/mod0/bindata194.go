package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHm, nil
}

func dictMod0IpaNeologdDictHm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
