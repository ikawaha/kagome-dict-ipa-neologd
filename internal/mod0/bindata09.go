package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAj, nil
}

func dictMod0IpaNeologdDictAj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.aj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
