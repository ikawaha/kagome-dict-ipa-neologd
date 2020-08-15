package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEo, nil
}

func dictMod0IpaNeologdDictEo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
