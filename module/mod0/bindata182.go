package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHa, nil
}

func dictMod0IpaNeologdDictHa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ha", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
