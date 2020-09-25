package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAa, nil
}

func dictMod0IpaNeologdDictAa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.aa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
