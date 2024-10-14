package cfg

import (
	"time"

	"ttgoer/log"
	"ttgoer/utils"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Cfg struct {
	Bot struct {
		Token          string        `mapstructure:"token" required:"true"`
		PollerTimeout  time.Duration `mapstructure:"poller_timeout" required:"true"`
		AdminId        int64         `mapstructure:"admin_id" required:"true"`
		AllowedUserIds []int64       `mapstructure:"allowed_user_ids"`
	} `mapstructure:"bot"`
	TikTok struct {
		DownloadTimeout    time.Duration `mapstructure:"download_timeout"`
		ShutdownTimeout    time.Duration `mapstructure:"shutdown_timeout"`
		DownloadMaxRetries uint          `mapstructure:"download_max_retries" required:"true"`
	} `mapstructure:"tik_tok"`
}

func (c *Cfg) IsUserAllowed(ID int64) bool {
	if c.Bot.AllowedUserIds == nil || c.Bot.AdminId == ID {
		return true
	}

	return utils.Contains(c.Bot.AllowedUserIds, ID)
}

var config = load()

func Get() *Cfg {
	if config == nil {
		log.S().Fatal("config not initialized")
	}
	return config
}

func load() *Cfg {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("tik_tok.download_timeout", 15*time.Second)
	viper.SetDefault("tik_tok.shutdown_timeout", 20*time.Second)

	if err := viper.ReadInConfig(); err != nil {
		log.S().Fatalf("error reading config file: %v", err)
	}

	var config Cfg
	if err := viper.Unmarshal(&config); err != nil {
		log.S().Fatalf("unable to unmarshal config: %v", err)
	}

	validateRequiredFields(config)

	yml := getAsYaml()
	log.S().Infof("loaded config:\n%s", yml)

	return &config
}

func getAsYaml() string {
	allSettings := viper.AllSettings()
	botSettings := allSettings["bot"].(map[string]any)
	botSettings["token"] = "<masked>"

	pretty, err := yaml.Marshal(allSettings)
	if err != nil {
		log.S().Fatalf("failed to marshal config to YAML: %v", err)
	}

	return string(pretty)
}
