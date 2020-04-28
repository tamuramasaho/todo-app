package image

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
)

// 使い方
// image.Upload(Entity.ImagePath(), c.FormFile("main_img"))
// 引数１つ目は、string型の画像のパス
// 引数２つ目はファイル echo.Context.FormFile("画像の名前")で取得できる
func Upload(path string, file *multipart.FileHeader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	wc := client.Bucket(os.Getenv("GCS_IMAGE_RAW_BUCKET")).Object(path).NewWriter(ctx)
	if _, err = io.Copy(wc, src); err != nil {
		return err
	}
	if err = wc.Close(); err != nil {
		return err
	}
	return nil
}