package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadImageImageResponseEntity struct{
    Url	string	`json:"url"`
    HashCode	string	`json:"hash_code"`
}
func (g UploadImageImageResponseEntity) String() string {
    return lib.ObjectToString(g)
}
