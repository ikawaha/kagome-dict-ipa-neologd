package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEi, nil
}

func dictMod0IpaNeologdDictEi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ei", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
