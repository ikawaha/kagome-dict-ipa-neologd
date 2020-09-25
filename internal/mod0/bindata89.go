package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDl, nil
}

func dictMod0IpaNeologdDictDl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
