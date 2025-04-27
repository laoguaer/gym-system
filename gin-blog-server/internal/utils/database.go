package utils

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
