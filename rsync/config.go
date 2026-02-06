package rsync

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/toastsandwich/rsync/utils"
	terr "github.com/toastsandwich/terror"
)

type SetConfigOptions struct {
	Alias    string
	Hostname string
	Username string
	Password string
}

type Config struct {
	Alias    string `json:"-"`
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Config) String() string {
	return fmt.Sprintf(
		`alias: %s
host: %s
username: %s
password: 
`, c.Alias, c.Hostname, c.Username)
}

func SetConfig(o SetConfigOptions) error {
	f, err := utils.OpenConfig()
	if err != nil {
		return terr.Wrap(err, OPEN_CONFIG_ERR_MSG)
	}
	defer f.Close()

	var configs = make(map[string]Config)
	err = json.NewDecoder(f).Decode(&configs)
	if err != nil && err != io.EOF {
		return terr.Wrap(err, "decoding configs")

	}

	config := Config{
		Alias:    o.Alias,
		Username: o.Username,
		Hostname: o.Hostname,
		Password: o.Password,
	}

	configs[o.Alias] = config

	encodedCfgs, err := json.MarshalIndent(&configs, "", " ")
	if err != nil {
		terr.Wrap(err, "indenting config json")
	}

	f.Seek(0, 0)
	f.Truncate(0)
	if _, err := f.Write(encodedCfgs); err != nil {
		return terr.Wrap(err, WRITE_CONFIG_ERR_MSG)
	}
	return nil
}

func GetConfig(alias string) (Config, error) {
	f, err := utils.OpenConfig()
	if err != nil {
		return Config{}, terr.Wrap(err, OPEN_CONFIG_ERR_MSG)
	}
	defer f.Close()

	configs := make(map[string]Config)
	if err := json.NewDecoder(f).Decode(&configs); err != nil {
		return Config{}, terr.Wrap(err, "decoding configs")
	}
	if config, ok := configs[alias]; ok {
		config.Alias = alias
		return config, nil
	}
	return Config{}, terr.Newf("config missing for %s", alias)
}

// func RemoveConfig(alias string) error {
// 	f, err := utils.OpenConfig()
// 	if err != nil {
// 		return terr.Wrap(err, OPEN_CONFIG_ERR_MSG)
// 	}
// 	defer f.Close()
// 	 	delete()
// }
