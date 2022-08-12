package aria2go

const DEFAULT_JSONRPC_VERSION = "2.0"
const DEFAULT_ARIA2_PORT = "6800"
const DEFAULT_ARIA2_ADDR = "127.0.0.1"
const DEFAULT_CONTENT_TYPE = "application/json"

type PositionOpt string

const (
	POS_SET PositionOpt = "POS_SET"
	POS_CUR PositionOpt = "POS_CUR"
	POS_END PositionOpt = "POS_END"
)