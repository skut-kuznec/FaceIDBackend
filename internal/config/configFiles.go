package config

import "github.com/spf13/viper"

type Files struct {
	filesUploadDir string
}

func (f *Files) ConfigFiles() error {
	f.filesUploadDir = viper.GetString("files.upload_dir")
	if f.filesUploadDir == "" {
		f.filesUploadDir = "upl"
	}
	return nil
}

func (f *Files) FileUploadDir() string {
	return f.filesUploadDir
}
