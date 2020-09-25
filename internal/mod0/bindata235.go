package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJb, nil
}

func dictMod0IpaNeologdDictJb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
