package consts

const (
	NoteTableName   = "note"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiServiceName  = "api"
	UserServiceName = "user"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
	Neo4jUri        = "neo4j://localhost:7687"
	NacosLogDir     = ""
	JWTSecretKey    = "chen"
)
