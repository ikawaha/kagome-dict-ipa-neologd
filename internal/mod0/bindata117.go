package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEn, nil
}

func dictMod0IpaNeologdDictEn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.en", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
