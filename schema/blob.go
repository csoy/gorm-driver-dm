package dmSchema

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/csoy/gorm-driver-dm/dmr"
)

type Blob string

func (blob Blob) Value() (driver.Value, error) {
	if len(blob) == 0 {
		return nil, nil
	}
	return string(blob), nil
}

func (blob *Blob) Scan(v interface{}) error {
	switch v.(type) {
	case *dmr.DmBlob:
		tmp := v.(*dmr.DmBlob)
		le, err := tmp.GetLength()
		if err != nil {
			return errors.New(fmt.Sprint("err：", err))
		}
		str := make([]byte, le)
		_, err = tmp.Read(str)
		*blob = Blob(str)
		break

	default:
		*blob = Blob(v.(string))
	}
	return nil
}
