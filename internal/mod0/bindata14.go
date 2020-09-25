package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAo, nil
}

func dictMod0IpaNeologdDictAo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ao", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
