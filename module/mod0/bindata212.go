package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIe, nil
}

func dictMod0IpaNeologdDictIe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ie", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
