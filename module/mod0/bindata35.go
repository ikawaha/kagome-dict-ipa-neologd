package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBj, nil
}

func dictMod0IpaNeologdDictBj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
