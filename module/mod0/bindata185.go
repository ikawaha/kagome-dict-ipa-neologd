package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHd, nil
}

func dictMod0IpaNeologdDictHd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
