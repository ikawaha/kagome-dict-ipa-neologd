package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEu, nil
}

func dictMod0IpaNeologdDictEu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
