package user

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/powehihihi/jobleaf/resume"
)

const (
	serviceName = "jobleaf"
	resumeyaml  = "resume.yaml"
)

type Config struct {
	resumePath string
}

func InitializeConfig() (*Config, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("os.UserConfigDir: %w", err)
	}

	if err = os.MkdirAll(configPath+"/"+serviceName, 0o644); err != nil {
		return nil, fmt.Errorf("os.MkdirAll(%s/%s: %w", configPath, serviceName, err)
	}

	resumePath := configPath + "/" + serviceName + "/" + resumeyaml

	_, err = os.Open(configPath + "/" + serviceName + "/" + resumeyaml)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(resumePath)
			if err != nil {
				return nil, fmt.Errorf("os.Create(%s): %w", resumePath, err)
			}
		}
		return nil, fmt.Errorf(
			"os.Open(%s/%s/%s) unexpected error: %w",
			configPath,
			serviceName,
			resumeyaml,
			err,
		)
	}

	return &Config{configPath + "/" + serviceName + "/" + resumeyaml}, nil
}

func (cfg *Config) LoadResume() (*resume.Resume, error) {
	var resume resume.Resume

	file, err := os.ReadFile(cfg.resumePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%s): %w", cfg.resumePath, err)
	}

	if err := yaml.Unmarshal(file, resume); err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal(%s, {}): %w", file, err)
	}

	return &resume, nil
}

func (cfg *Config) SaveResume(resume *resume.Resume) error {
	yaml, err := yaml.Marshal(resume)
	if err != nil {
		return fmt.Errorf("marshalling error: %w", err)
	}
	if err := os.WriteFile(cfg.Path(), yaml, 0o644); err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}

	return nil
}

func (cfg *Config) Path() string {
	return cfg.resumePath
}
