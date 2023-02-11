package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/publish"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/minio"
	"github.com/gofrs/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type ActionPulishService struct {
	ctx     context.Context
	session neo4j.SessionWithContext
}

func NewActionPulishService(ctx context.Context, driver neo4j.DriverWithContext) *ActionPulishService {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	return &ActionPulishService{ctx: ctx, session: session}
}

func (s *ActionPulishService) ActionPulish(req *publish.ActionRequest) error {
	MinioVideoBucketName := consts.MinioVideoBucketName
	videoData := []byte(req.Data)

	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"

	// Upload video file
	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}

	// Get video uri
	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}

	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := captureCover(playUrl)
	if err != nil {
		return err
	}

	// Upload video cover file
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}

	// Get video cover uri
	coverUrl, err := minio.GetFileUrl(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	_ = strings.Split(coverUrl.String(), "?")[0]

	return nil
}

// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func captureCover(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	return buf.Bytes(), err
}
