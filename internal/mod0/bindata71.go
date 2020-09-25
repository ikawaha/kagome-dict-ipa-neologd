package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCt, nil
}

func dictMod0IpaNeologdDictCt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ct", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
