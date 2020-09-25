package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHf, nil
}

func dictMod0IpaNeologdDictHf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
