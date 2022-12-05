package mediacenter
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	mediacenterentity "github.com/wjpxxx/lazadago/mediacenter/entity"
)
//MediaCenter
type MediaCenter struct{
	Config *lazadaConfig.Config
}

//CompleteCreateVideo
//@Title After uploading all blocks of the video file,  call CompleteCreateVideo to complete the video uploading process.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/block/commit
func (s *MediaCenter) CompleteCreateVideo (uploadId string,parts string,title string,coverUrl string) mediacenterentity.CompleteCreateVideoResult {
    method := "/media/video/block/commit"
    params := lib.InRow{
      "uploadId":uploadId,
      "parts":parts,
      "title":title,
      "coverUrl":coverUrl,
    }
    result := mediacenterentity.CompleteCreateVideoResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetVideo
//@Title You call this action to get video info after uploading.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/get
func (s *MediaCenter) GetVideo (videoId int64) mediacenterentity.GetVideoResult {
    method := "/media/video/get"
    params := lib.InRow{
      "videoId":videoId,
    }
    result := mediacenterentity.GetVideoResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetVideoQuota
//@Title You call this api to get the capacity quota of seller.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/quota/get
func (s *MediaCenter) GetVideoQuota () mediacenterentity.GetVideoQuotaResult {
    method := "/media/video/quota/get"
    params := lib.InRow{}
    result := mediacenterentity.GetVideoQuotaResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//InitCreateVideo
//@Title A seller starts to upload a video file
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/block/create
func (s *MediaCenter) InitCreateVideo (fileName string,fileBytes int) mediacenterentity.InitCreateVideoResult {
    method := "/media/video/block/create"
    params := lib.InRow{
      "fileName":fileName,
      "fileBytes":fileBytes,
    }
    result := mediacenterentity.InitCreateVideoResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RemoveVideo
//@Title You can this api to delete a video file permanently.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/remove
func (s *MediaCenter) RemoveVideo (videoId int64) mediacenterentity.RemoveVideoResult {
    method := "/media/video/remove"
    params := lib.InRow{
      "videoId":videoId,
    }
    result := mediacenterentity.RemoveVideoResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadVideoBlock
//@Title The API is used to upload one block of origin video file. The video file can split into multiple files. For example, a 8MB video file can be split into three blocks. 3MB, 3MB and 2MB. These three blocks can be uploaded by calling UploadVideoBlock three times.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=32&path=/media/video/block/upload
func (s *MediaCenter) UploadVideoBlock (uploadId string,blockNo string,blockCount string,file []byte) mediacenterentity.UploadVideoBlockResult {
    method := "/media/video/block/upload"
    params := lib.InRow{
      "uploadId":uploadId,
      "blockNo":blockNo,
      "blockCount":blockCount,
      "@file":file,
    }
    result := mediacenterentity.UploadVideoBlockResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
