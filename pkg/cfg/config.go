package cfg

import "path"

// Config holds bot config information
type Config struct {
	configDir string

	Wadl struct {
		LogLevel string `toml:"log-level"`
		Prefix   string `toml:"prefix"`
		Token    string `toml:"token"`
		GuildID  string `toml:"guild-id"`
	} `toml:"waddles" comment:"General Bot Configuration"`

	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		Name string `toml:"database-name"`
		URL  string `toml:"url" commented:"true" comment:"uncomment to use a postgres URI instead of above"`
	} `toml:"database" comment:"Postgresql Database Connection Information"`

	NitroPerk struct {
		BoosterChannel struct {
			ParentID string `toml:"parent-id" comment:"Discord catagory ID for channels to be managed under"`
		} `toml:"booster-channel" comment:"server booster personal channel options"`
	} `toml:"nitro" comment:"perks related to being a server booster"`

	Gatekeeper struct {
		Enabled    bool   `toml:"enabled"`
		ChannelID  string `toml:"channel-id"`
		RoleID     string `toml:"role-id"`
		WelcomeMsg string `toml:"welcome-msg"`
		Rules      string `toml:"rules"`
	} `toml:"gatekeeper" comment:"settings related to the gatekeeper function"`

	Permissions struct {
		DebugUsers []string `toml:"bypass-users" comment:"list of user IDs who bypass the permission's system (useful for testing)"` //TODO: reorganize debugusers to different struct
	} `toml:"permissions" comment:"settings related to the permissions system"`
}

//GetConfigFileLocation returns the full path of the requested configFile
func (config Config) GetConfigFileLocation(configFile string) string {
	return path.Clean(config.configDir + "/" + configFile)
}
