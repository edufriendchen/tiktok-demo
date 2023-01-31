package consts

const (
	ApiServiceName      = "api"
	ApiServiceAddr      = ":1060"
	UserServiceName     = "user"
	UserServiceAddr     = ":9000"
	RelationServiceName = "relation"
	RelationServiceAddr = ":9100"
	MessageServiceName  = "message"
	MessageServiceAddr  = ":9200"
	CommentServiceName  = "comment"
	CommentServiceAddr  = ":9300"
	FavoriteServiceName = "favorite"
	FavoriteServiceAddr = ":9400"
	FeedServiceName     = "feed"
	FeedServiceddr      = ":9500"
	PushlishServiceName = "puslish"
	PushlishServiceddr  = ":9600"
)

const (
	NoteTableName   = "note"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	ExportEndpoint  = ":4317"
	DefaultLimit    = 10
	Neo4jUri        = "neo4j://localhost:7687"
	Neo4jUsername   = "neo4j"
	Neo4jPassword   = "friendchen0429"
	NacosAddr       = "127.0.0.1"
	NacosPoint      = 8848
	NacosLogDir     = ""
	JWTSecretKey    = "chen"
)
