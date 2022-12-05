package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type MigrateImagesResult struct{
    BatchId	string	`json:"batch_id"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]MigrateImagesDetailResponseEntity	`json:"detail"`
}
func (g MigrateImagesResult) String() string {
    return lib.ObjectToString(g)
}

type MigrateImagesDetailResponseEntity struct{}