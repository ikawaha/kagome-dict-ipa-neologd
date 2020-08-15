package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEz, nil
}

func dictMod0IpaNeologdDictEz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ez", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
