package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEa, nil
}

func dictMod0IpaNeologdDictEa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ea", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
