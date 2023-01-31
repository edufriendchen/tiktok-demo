package global

import (
	"github.com/edufriendchen/tiktok-demo/pkg/jwt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	Neo4jSession neo4j.SessionWithContext
	Jwt          *jwt.JWT
)
