package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// type state struct {
// 	ConfigPTR *Config
// }

type Config struct {
	DBURL    string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	var config Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("failed to get config file path: %w", err)
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return config, nil

}

func (cfg *Config) SetUser(username string) error {

	cfg.UserName = username

	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user's home directory: %w", err)
	}

	configFilePath := filepath.Join(userHomeDirectory, configFileName)

	return configFilePath, nil
}

func write(cfg Config) error {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config struct: %w", err)
	}

	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}

// type command struct {
// 	name string
// 	args []string
// }

// type commands struct {
// 	commandMap map[string]func(*state, command) error
// }

// func handlerLogin(s *state, cmd command) error {

// 	if len(cmd.args) == 0 {
// 		return errors.New("the login handler expects a single argument, the username")
// 	}

// 	username := cmd.args[0]

// 	s.ConfigPTR.SetUser(username)

// 	fmt.Printf("%s has been set as the current username", username)

// 	return nil
// }

// func (c *commands) run(s *state, cmd command) error {

// 	if command, exists := c.commandMap[cmd.name]; exists {
// 		err := command(s, cmd)
// 		if err != nil {
// 			return fmt.Errorf("unable to run command %w", err)
// 		}
// 	} else {
// 		return errors.New("command doesn't exist")
// 	}

// 	return nil
// }

// func (c *commands) register(name string, f func(*state, command) error) {
// 	c.commandMap[name] = f
// }

// func StartGator() {

// 	cfg, err := Read()
// 	if err != nil {
// 		log.Fatalf("Critical error opening config: %v", err)
// 	}

// 	myState := state{
// 		ConfigPTR: &cfg,
// 	}

// 	myCommands := commands{
// 		commandMap: make(map[string]func(*state, command) error),
// 	}

// 	myCommands.register("login", handlerLogin)

// 	userArgsCount := len(os.Args)

// 	if userArgsCount < 2 {
// 		log.Fatalln("Error occurred: not enough arguments")
// 	}

// 	myCommand := command{
// 		name: os.Args[1],
// 		args: os.Args[2:],
// 	}

// 	err = myCommands.run(&myState, myCommand)
// 	if err != nil {
// 		log.Fatal("Error occurred:", err)
// 	}
// }
