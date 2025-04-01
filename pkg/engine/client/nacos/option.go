package nacos

import "log"

// Config ...
type Config struct {
	TimeoutMs            uint64   `yaml:"timeoutMs"` // timeout for requesting Nacos server, default value is 10000ms
	ListenInterval       uint64   // Deprecated
	BeatInterval         int64    // the time interval for sending beat to server,default value is 5000ms
	NamespaceId          string   `yaml:"namespaceId"` // the namespaceId of Nacos.When namespace is public, fill in the blank string here.
	AppName              string   // the appName
	Endpoints            []string `yaml:"endpoints"` // the endpoint for get Nacos server addresses
	RegionId             string   // the regionId for kms
	AccessKey            string   // the AccessKey for kms
	SecretKey            string   // the SecretKey for kms
	OpenKMS              bool     // it's to open kms,default is false. https://help.aliyun.com/product/28933.html
	CacheDir             string   `yaml:"cacheDir"` // the directory for persist nacos service info,default value is current path
	UpdateThreadNum      int      // the number of gorutine for update nacos service info,default value is 20
	NotLoadCacheAtStart  bool     `yaml:"loadCacheAtStart"` // not to load persistent nacos service info in CacheDir at start time
	UpdateCacheWhenEmpty bool     // update cache when get empty service instance from server
	Username             string   `yaml:"username"` // the username for nacos auth
	Password             string   `yaml:"password"` // the password for nacos auth
	ContextPath          string   // the nacos server contextpath
	AppendToStdout       bool     // append log to stdout
	EnableTrace          bool
	Name                 string
	logger               *log.Logger
}
