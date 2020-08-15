package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHo, nil
}

func dictMod0IpaNeologdDictHo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ho", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
