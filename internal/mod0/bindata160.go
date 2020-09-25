package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGe, nil
}

func dictMod0IpaNeologdDictGe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ge", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
