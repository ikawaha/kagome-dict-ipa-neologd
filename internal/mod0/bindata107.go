package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEd, nil
}

func dictMod0IpaNeologdDictEd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ed", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
