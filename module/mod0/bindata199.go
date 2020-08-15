package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHr, nil
}

func dictMod0IpaNeologdDictHr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
