package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	DEFAULT_DEBUG = false
	DEFAULT_RPS   = 7

	DEFAULT_URL          = "https://indexer.dydx.trade/v4"
	DEFAULT_CANDLES_PATH = "/candles/perpetualMarkets"
	DEFAULT_MARKETS_PATH = "/perpetualMarkets"
)

type Config struct {
	Debug bool  `koanf:"debug"`
	Rps   uint8 `koanf:"rps"`

	Url         string `koanf:"url"`
	CandlesPath string `koanf:"candles_path"`
	MarketsPath string `koanf:"markets_path"`
}

func MustNew() *Config {
	var c Config

	k := koanf.New(".")

	mustLoadDefaults(k)

	fileFlag := mustCheckFileFlag()
	if fileFlag != "" {
		mustLoadYamlFile(k, fileFlag)
	}

	mustLoadEnv(k)

	err := k.Unmarshal("", &c)
	if err != nil {
		panic(fmt.Errorf("error while unmarshalling config: %w", err))
	}

	return &c
}

func mustLoadDefaults(k *koanf.Koanf) {
	err := k.Load(confmap.Provider(map[string]interface{}{
		"debug": DEFAULT_DEBUG,

		"rps":          DEFAULT_RPS,
		"url":          DEFAULT_URL,
		"candles_path": DEFAULT_CANDLES_PATH,
		"markets_path": DEFAULT_MARKETS_PATH,
	}, "."), nil)
	if err != nil {
		panic(fmt.Errorf("error while loading config defaults: %w", err))
	}
}

func mustCheckFileFlag() string {
	var fFlag = flag.String("ff", "", "Path to the configuration YAML file")

	flag.Parse()

	return *fFlag
}

func mustLoadYamlFile(k *koanf.Koanf, name string) {
	err := k.Load(file.Provider(name), yaml.Parser())
	if err != nil {
		panic(fmt.Errorf("error while loading yaml config file: %w", err))
	}
}

func mustLoadEnv(k *koanf.Koanf) {
	err := k.Load(env.Provider("DYDXV4_", ".", func(s string) string {
		return strings.Replace(strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "DYDXV4_")), "_", ".", -1), "-", "_", -1)
	}), nil)

	if err != nil {
		panic(err)
	}
}
