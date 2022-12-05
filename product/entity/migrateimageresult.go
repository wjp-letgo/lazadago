package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type MigrateImageResult struct{
    Data	MigrateImageDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]MigrateImageDetailResponseEntity	`json:"detail"`
}
func (g MigrateImageResult) String() string {
    return lib.ObjectToString(g)
}
type MigrateImageDetailResponseEntity struct{}