package helper

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func UploadAndCompressImage(file *multipart.FileHeader, maxSizeKB uint) (string, error) {
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Validasi ekstensi
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("invalid file format")
	}

	// Decode gambar
	var img image.Image
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(src)
	case ".png":
		img, err = png.Decode(src)
	}
	if err != nil {
		return "", err
	}

	// Kompresi
	compressedImg := resize.Thumbnail(800, 800, img, resize.Lanczos3)

	// Simpan gambar
	shortUUID := uuid.New().String()[:8]
	fileName := shortUUID + ext
	filePath := filepath.Join("uploads", fileName)
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	switch ext {
	case ".jpg", ".jpeg":
		jpeg.Encode(out, compressedImg, &jpeg.Options{Quality: 70})
	case ".png":
		png.Encode(out, compressedImg)
	}

	// Cek ukuran file
	if info, _ := out.Stat(); info.Size() > int64(maxSizeKB*1024) {
		os.Remove(filePath)
		return "", fmt.Errorf("image size exceeds limit")
	}

	return filePath, nil
}
