package aria2go

// Option 请求参数配置
// 参数参考如下官网配置说明
// http://aria2.github.io/manual/en/html/aria2c.html#input-file
type Option struct {
	Dir                           string `json:"dir"`
	Out                           string `json:"out"`
	Gid                           string `json:"gid"`
	AllProxy                      string `json:"all-proxy"`
	AllProxyPasswd                string `json:"all-proxy-passwd"`
	AllProxyUser                  string `json:"all-proxy-user"`
	AllowOverwrite                string `json:"allow-overwrite"`
	AllowPieceLengthChange        string `json:"allow-piece-length-change"`
	AlwaysResume                  string `json:"always-resume"`
	AsyncDns                      string `json:"async-dns"`
	AutoFileRenaming              string `json:"auto-file-renaming"`
	BTEnableHookAfterHashCheck    string `json:"bt-enable-hook-after-hash-check"`
	BTEnableLpd                   string `json:"bt-enable-Lpd"`
	BTExcludeTracker              string `json:"bt-exclude-tracker"`
	BTExternalIp                  string `json:"bt-external-ip"`
	BTForceEncryption             string `json:"bt-force-encryption"`
	BTHashCheckSeed               string `json:"bt-hash-check-seed"`
	BTLoadSavedMetadata           string `json:"bt-load-saved-metadata"`
	BTMaxPeers                    string `json:"bt-max-peers"`
	BTMetadataOnly                string `json:"bt-metadata-only"`
	BTMinCryptoLevel              string `json:"bt-min-crypto-level"`
	BTPrioritizePiece             string `json:"bt-prioritize-piece"`
	BTRemoveUnselectedFile        string `json:"bt-remove-unselected-file"`
	BTRequestPeerSpeedLimit       string `json:"bt-request-peer-speed-limit"`
	BTRequireCrypto               string `json:"bt-require-crypto"`
	BTSaveMetadata                string `json:"bt-save-metadata"`
	BTSeedUnverified              string `json:"bt-seed-unverified"`
	BTStopTimeout                 string `json:"bt-stop-timeout"`
	BTTracker                     string `json:"bt-tracker"`
	BTTrackerConnectTimeout       string `json:"bt-tracker-connect-timeout"`
	BTTrackerInterval             string `json:"bt-tracker-interval"`
	BTTrackerTimeout              string `json:"bt-tracker-timeout"`
	CheckIntegrity                string `json:"check-integrity"`
	Checksum                      string `json:"checksum"`
	ConditionalGet                string `json:"conditional-get"`
	ConnectTimeout                string `json:"connect-timeout"`
	ContentDispositionDefaultUtf8 string `json:"content-disposition-default-utf8"`
	Continue                      string `json:"continue"`
	DryRun                        string `json:"dry-run"`
	EnableHTTPKeepAlive           string `json:"enable-http-keep-alive"`
	EnableHTTPPipelining          string `json:"enable-http-pipelining"`
	EnableMmap                    string `json:"enable-mmap"`
	EnablePeerExchange            string `json:"enable-peer-exchange"`
	FileAllocation                string `json:"file-allocation"`
	FollowMetalink                string `json:"follow-metalink"`
	FollowTorrent                 string `json:"follow-torrent"`
	ForceSave                     string `json:"force-save"`
	FTPPasswd                     string `json:"ftp-passwd"`
	FTPPasv                       string `json:"ftp-pasv"`
	FTPProxy                      string `json:"ftp-proxy"`
	FTPProxyPasswd                string `json:"ftp-proxy-passwd"`
	FTPProxyUser                  string `json:"ftp-proxy-user"`
	FTPReuseConnection            string `json:"ftp-reuse-connection"`
	FTPType                       string `json:"ftp-type"`
	FTPUser                       string `json:"ftp-user"`
	HashCheckOnly                 string `json:"hash-check-only"`
	Header                        string `json:"header"`
	HTTPAcceptGzip                string `json:"http-accept-gzip"`
	HTTPAuthChallenge             string `json:"http-auth-challenge"`
	HTTPNoCache                   string `json:"http-no-cache"`
	HTTPPasswd                    string `json:"http-passwd"`
	HTTPProxy                     string `json:"http-proxy"`
	HTTPProxyPasswd               string `json:"http-proxy-passwd"`
	HTTPProxyUser                 string `json:"http-proxy-user"`
	HTTPUser                      string `json:"http-user"`
	HTTPSProxy                    string `json:"https-proxy"`
	HTTPSProxyPasswd              string `json:"https-proxy-passwd"`
	HTTPSProxyUser                string `json:"https-proxy-user"`
	IndexOut                      string `json:"index-out"`
	LowestSpeedLimit              string `json:"lowest-speed-limit"`
	MaxConnectionPerServer        string `json:"max-connection-per-server"`
	MaxDownloadLimit              string `json:"max-download-limit"`
	MaxFileNotFound               string `json:"max-file-not-found"`
	MaxMmapLimit                  string `json:"max-mmap-limit"`
	MaxResumeFailureTries         string `json:"max-resume-failure-tries"`
	MaxTries                      string `json:"max-tries"`
	MaxUploadLimit                string `json:"max-upload-limit"`
	MetalinkBaseUri               string `json:"metalink-base-uri"`
	MetalinkEnableUniqueProtocol  string `json:"metalink-enable-unique-protocol"`
	MetalinkLanguage              string `json:"metalink-language"`
	MetalinkLocation              string `json:"metalink-location"`
	MetalinkOs                    string `json:"metalink-os"`
	MetalinkPreferredProtocol     string `json:"metalink-preferred-protocol"`
	MetalinkVersion               string `json:"metalink-version"`
	MinSplitSize                  string `json:"min-split-size"`
	NoFileAllocationLimit         string `json:"no-file-allocation-limit"`
	NoNetrc                       string `json:"no-netrc"`
	NoProxy                       string `json:"no-proxy"`
	ParameterizedUri              string `json:"parameterized-uri"`
	Pause                         string `json:"pause"`
	PauseMetadata                 string `json:"pause-metadata"`
	PieceLength                   string `json:"piece-length"`
	ProxyMethod                   string `json:"proxy-method"`
	RealtimeChunkChecksum         string `json:"realtime-chunk-checksum"`
	Referer                       string `json:"referer"`
	RemoteTime                    string `json:"remote-time"`
	RemoveControlFile             string `json:"remove-control-file"`
	RetryWait                     string `json:"retry-wait"`
	ReuseUri                      string `json:"reuse-uri"`
	RPCSaveUploadMetadata         string `json:"rpc-save-upload-metadata"`
	SeedRatio                     string `json:"seed-ratio"`
	SeedTime                      string `json:"seed-time"`
	SelectFile                    string `json:"select-file"`
	Split                         string `json:"split"`
	SSHHostKeyMD                  string `json:"ssh-host-key-md"`
	StreamPieceSelector           string `json:"stream-piece-selector"`
	Timeout                       string `json:"timeout"`
	UriSelector                   string `json:"uri-selector"`
	UseHead                       string `json:"use-head"`
	UserAgent                     string `json:"user-agent"`
	Position                      string `json:"position"`
}

// Response aria2 通常响应
type Response struct {
	ID      string         `json:"id"`
	JSONRPC string         `json:"jsonrpc"`
	Result  string         `json:"result"`
	Error   *ResponseError `json:"error"`
}

// ResponseError Response 中 Error 字段
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
