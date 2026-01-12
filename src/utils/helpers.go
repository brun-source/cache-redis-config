package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-redis/redis/v8"
)

func getRedisConfig() (*redis.Options, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	if redisHost == "" || redisPort == "" {
		return nil, errors.New("redis host and port are required")
	}

	config := &redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       getRedisDB(redisDB),
	}

	return config, nil
}

func getRedisDB(db string) int {
	if db == "" {
		return 0
	}

	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Printf("invalid redis db: %s", db)
		return 0
	}

	return dbInt
}

func loadConfigFromFile(filePath string, config interface{}) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		return err
	}

	return nil
}

func getAbsolutePath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(workingDir, path), nil
}

func isValidRedisKey(key string) bool {
	return strings.TrimSpace(key) != ""
}