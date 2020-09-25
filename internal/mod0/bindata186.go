package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHe, nil
}

func dictMod0IpaNeologdDictHe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.he", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
