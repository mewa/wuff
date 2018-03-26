package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Email   string
	Smtp    *Smtp
	Service []*Service
}

type Service struct {
	Name        string
	Retries     int
	RetryPeriod int
	CheckPeriod int
	Check       string
	Start       string
}

type Smtp struct {
	Server   string
	Sender   string
	Port     int
	User     string
	Password string
}

func ReadConfig() (*Config, error) {
	var conf Config

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error in config file: %s", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("Invalid config structure: %s", err)
	}

	return &conf, nil
}

func (conf *Config) Verify() error {
	if conf.Email == "" {
		return fmt.Errorf("Email not supplied")
	}

	err := conf.Smtp.Verify()

	if err != nil {
		return err
	}

	for _, service := range conf.Service {
		err = service.Verify()

		if err != nil {
			return err
		}
	}

	return nil
}

func (smtp *Smtp) Verify() error {
	if smtp.Server == "" {
		return fmt.Errorf("SMTP server not provided")
	}

	if smtp.User == "" {
		return fmt.Errorf("SMTP server user not provided")
	}

	if smtp.Password == "" {
		return fmt.Errorf("SMTP server password not provided")
	}

	if smtp.Port == 0 {
		smtp.Port = 587
	}

	return nil
}

func (s *Service) SetDefaultCheck() {
	s.Check = fmt.Sprintf("service %s status", s.Name)
}

func (s *Service) SetDefaultStart() {
	s.Start = fmt.Sprintf("service %s start", s.Name)
}

func (service *Service) Verify() error {
	if service.Name == "" {
		return fmt.Errorf("Service must include name")
	}

	if service.CheckPeriod == 0 {
		service.CheckPeriod = 10
	}

	if service.Check == "" {
		service.SetDefaultCheck()
	}

	if service.Start == "" {
		service.SetDefaultStart()
	}

	return nil
}
