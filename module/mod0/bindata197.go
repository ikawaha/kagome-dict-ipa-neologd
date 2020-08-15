package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHp, nil
}

func dictMod0IpaNeologdDictHp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
