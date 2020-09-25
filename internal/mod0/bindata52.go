package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCa, nil
}

func dictMod0IpaNeologdDictCa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ca", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
