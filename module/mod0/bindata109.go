package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEf, nil
}

func dictMod0IpaNeologdDictEf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ef", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
