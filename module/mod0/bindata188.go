package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHg, nil
}

func dictMod0IpaNeologdDictHg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
