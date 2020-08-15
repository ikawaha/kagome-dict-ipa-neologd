package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAz, nil
}

func dictMod0IpaNeologdDictAz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.az", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
