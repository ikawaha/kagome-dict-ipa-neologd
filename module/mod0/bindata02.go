package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAc, nil
}

func dictMod0IpaNeologdDictAc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ac", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
