package global

import (
	"github.com/edufriendchen/tiktok-demo/pkg/jwt"
	"github.com/minio/minio-go/v7"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gorm.io/gorm"
)

var (
	Neo4jDriver neo4j.DriverWithContext
	Jwt         *jwt.JWT
	DB          *gorm.DB
	MinioClient *minio.Client
)
