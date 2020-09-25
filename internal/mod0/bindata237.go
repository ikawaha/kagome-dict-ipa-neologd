package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJd, nil
}

func dictMod0IpaNeologdDictJd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
