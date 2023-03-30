package config

import "github.com/spf13/viper"

type Files struct {
	imageUploadDir string
}

func (f *Files) ConfigImage() error {
	f.imageUploadDir = viper.GetString("image.upload_dir")
	if f.imageUploadDir == "" {
		f.imageUploadDir = "upl"
	}
	return nil
}

func (f *Files) ImageUploadDir() string {
	return f.imageUploadDir
}
