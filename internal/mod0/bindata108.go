package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEe, nil
}

func dictMod0IpaNeologdDictEe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ee", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
