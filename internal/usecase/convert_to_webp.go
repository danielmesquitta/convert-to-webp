package usecase

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nickalie/go-webpbin"
)

func ConvertToWebp(path string, quality uint) error {
	if quality > 100 {
		quality = 80
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	processFile := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			log.Println("Processing file:", path)
			if err := convertFileToWebp(path, quality); err != nil {
				log.Printf("Failed to convert %s: %v", path, err)
			}
		}
		return nil
	}

	if info.IsDir() {
		if err := filepath.Walk(path, processFile); err != nil {
			return err
		}
		return nil
	}

	if err := processFile(path, info, nil); err != nil {
		return err
	}
	return nil
}

func convertFileToWebp(filePath string, quality uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".png":
		img, err = png.Decode(file)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	default:
		return fmt.Errorf("Unsupported file type %s", ext)
	}
	if err != nil {
		return err
	}

	out := strings.TrimSuffix(filePath, ext) + ".webp"
	cwebp := webpbin.NewCWebP()

	err = cwebp.InputImage(img).OutputFile(out).Quality(quality).Run()
	if err != nil {
		return err
	}

	log.Println("Converted to", out)
	return nil
}
