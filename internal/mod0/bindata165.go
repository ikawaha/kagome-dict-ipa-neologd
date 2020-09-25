package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGj, nil
}

func dictMod0IpaNeologdDictGj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
