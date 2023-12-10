package user

import (
	"fmt"
	"os"

	"github.com/powehihihi/jobleaf/internal/lib"
	"github.com/powehihihi/jobleaf/internal/model"
)

type Config struct {
	configDir      string
	userResumePath string
}

// Init initializes user configurations with content.
// If userResume already exists it's safe to pass nil content.
func Init(content []byte) (*Config, error) {
	cfg := &Config{}

	if err := cfg.InitConfigDir(); err != nil {
		return nil, fmt.Errorf("InitConfigDir: %w", err)
	}

	if err := cfg.InitUserResumePath(content); err != nil {
		return nil, fmt.Errorf("InitUserResumePath: %w", err)
	}

	return cfg, nil
}

// InitConfigDir creates user config directory.
// On success it initializes configDir.
func (cfg *Config) InitConfigDir() error {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("os.UserConfigDir: %w", err)
	}

	configDir := configPath + "/jobleaf"

	if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("os.MkdirAll(%s): %w", cfg.configDir, err)
	}

	cfg.configDir = configDir

	return nil
}

// InitUserResumePath works with resume.yaml file in user config directory.
// If not exists, it's created with content.
// On success (no errors) it initializes userResumePath.
// If userResume already exists it's safe to pass nil content.
func (cfg *Config) InitUserResumePath(content []byte) error {
	if cfg.configDir == "" {
		return errConfigDirNotInited
	}

	userResumePath := cfg.configDir + "/resume.yaml"

	if _, err := os.Open(userResumePath); err != nil {
		if os.IsNotExist(err) {
			if content == nil {
				return errPassedContentIsNil
			}

			if err = os.WriteFile(userResumePath, content, os.ModePerm); err != nil {
				return fmt.Errorf("os.WriteFile: %w", err)
			}

			cfg.userResumePath = userResumePath

			return nil
		}

		return fmt.Errorf("os.Open(%s) unexpected error: %w", cfg.userResumePath, err)
	}

	cfg.userResumePath = userResumePath

	return nil
}

func (cfg *Config) ConfigDir() string {
	return cfg.configDir
}

func (cfg *Config) UserResumePath() string {
	return cfg.userResumePath
}

// LoadResume opens file and unmarshalles it to model.Resume.
func (cfg *Config) LoadResume() (*model.Resume, error) {
	resume, err := lib.LoadFromFile(cfg.userResumePath)
	if err != nil {
		return nil, fmt.Errorf("LoadFromFile(%s)", cfg.userResumePath)
	}

	return resume, nil
}

// SaveResume saves resume.yaml to config directory.
func (cfg *Config) SaveResume(resume *model.Resume) error {
	if err := lib.SaveToFile(cfg.userResumePath, resume); err != nil {
		return fmt.Errorf("SaveToFile: %w", err)
	}

	return nil
}
