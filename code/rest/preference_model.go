package rest

import (
	"github.com/eyebluecn/tank/code/core"
	jsoniter "github.com/json-iterator/go"
)

type Preference struct {
	Base
	Name                  string `json:"name" gorm:"type:varchar(45)"`
	LogoUrl               string `json:"logoUrl" gorm:"type:varchar(255)"`
	FaviconUrl            string `json:"faviconUrl" gorm:"type:varchar(255)"`
	Copyright             string `json:"copyright" gorm:"type:varchar(1024)"`
	Record                string `json:"record" gorm:"type:varchar(1024)"`
	DownloadDirMaxSize    int64  `json:"downloadDirMaxSize" gorm:"type:bigint(20) not null;default:-1"`
	DownloadDirMaxNum     int64  `json:"downloadDirMaxNum" gorm:"type:bigint(20) not null;default:-1"`
	DefaultTotalSizeLimit int64  `json:"defaultTotalSizeLimit" gorm:"type:bigint(20) not null;default:-1"`
	AllowRegister         bool   `json:"allowRegister" gorm:"type:tinyint(1) not null;default:0"`
	PreviewConfig         string `json:"previewConfig" gorm:"type:text"`
	ScanConfig            string `json:"scanConfig" gorm:"type:text"`
	Version               string `json:"version" gorm:"-"`
}

// set File's table name to be `profiles`
func (this *Preference) TableName() string {
	return core.TABLE_PREFIX + "preference"
}

const (
	//scan scope all.
	SCAN_SCOPE_ALL = "ALL"
	//scan scope custom.
	SCAN_SCOPE_CUSTOM = "CUSTOM"
)

//scan config struct.
type ScanConfig struct {
	//whether enable the scan task.
	Enable bool `json:"enable"`
	//when to process the task.
	Cron string `json:"cron"`
	//username
	Usernames []string `json:"usernames"`
	//scan scope. see SCAN_SCOPE
	Scope string `json:"scope"`
}

//fetch the scan config
func (this *Preference) FetchScanConfig() *ScanConfig {

	json := this.ScanConfig
	if json == "" || json == EMPTY_JSON_MAP {

		return &ScanConfig{
			Enable: false,
		}
	} else {
		m := &ScanConfig{}

		err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(json), &m)
		if err != nil {
			panic(err)
		}
		return m
	}
}
