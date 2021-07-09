// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package config

// ProducerCfg represents the producer config.
type ProducerCfg struct {
	Nsqd  string `mapstructure:"nsqd"`
	Topic string `mapstructure:"topic"`
}

// ConsumerCfg represents the consumer config.
type ConsumerCfg struct {
	Lookupds    []string `mapstructure:"lookupds"`
	Topic       string   `mapstructure:"topic"`
	Channel     string   `mapstructure:"channel"`
	Concurrency int      `mapstructure:"concurrency"`
}

// AWSS3Cfg represents AWS S3 credentials.
type AWSS3Cfg struct {
	Region    string `mapstructure:"region"`
	SecretKey string `mapstructure:"secret_key"`
	AccessKey string `mapstructure:"access_key"`
}

// LocalFsCfg represents local file system storage data.
type LocalFsCfg struct {
	RootDir string `mapstructure:"root_dir"`
}

// StorageCfg represents the object storage config.
type StorageCfg struct {
	SharedVolume string     `mapstructure:"shared_volume"`
	Bucket       string     `mapstructure:"bucket"`
	Timeout      int        `mapstructure:"timeout"`
	S3           AWSS3Cfg   `mapstructure:"s3"`
	Local        LocalFsCfg `mapstructure:"local"`
}
