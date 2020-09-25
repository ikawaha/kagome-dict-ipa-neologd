package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEyBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEy, nil
}

func dictMod0IpaNeologdDictEy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ey", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
