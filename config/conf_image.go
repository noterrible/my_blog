package config

type Image struct {
	//Path         string `yaml:"path"`
	//enableOrigin    string `yaml:"enable_origin"`
	Size            int64  `yaml:"size"`
	BucketName      string `yaml:"bucket_name"`
	UploadDomain    string `yaml:"upload_domain"`
	GetDomain       string `yaml:"get_domain"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}
