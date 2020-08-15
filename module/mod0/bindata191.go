package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHj, nil
}

func dictMod0IpaNeologdDictHj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
