package amiga

import (
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/EventBus"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config is the interface that wraps the basic methods to get configuration values.
type Config interface {
	String(key string) (string, error)

	Strings(key string) ([]string, error)

	Int(key string) (int, error)

	Ints(key string) ([]int, error)

	Bool(key string) (bool, error)

	Bools(key string) ([]bool, error)

	Float64(key string) (float64, error)

	Float64s(key string) ([]float64, error)

	// Binds a struct to be automatically filled with configuration values and internally registered to be hot-reloaded on config-now file refresh
	Bind(c AmigaConfig) error

	// Return a ConfStruct with all the Amiga Fwk properties already bound to be hot-reloaded
	AmigaFwkConfig() *AmigaFwkConfig
}

// All structs that want to be automatically filled with configuration values must implement this interface
type AmigaConfig interface {
	Prefix() string
}

// Internal implementation of Config interface using Koanf library
type defaultConfig struct {
	koanf             *koanf.Koanf
	eventBus          EventBus.Bus
	boundStructs      []*AmigaConfig
	watcherRegistered bool
}

// Struct that maps all Amiga Fwk properties (TODO extract from here to Amiga core instance?)
var amigaFwkConfig = &AmigaFwkConfig{}

func NewConfig(eventBus EventBus.Bus) (Config, error) {

	c := &defaultConfig{
		eventBus: eventBus,
	}

	err := c.reloadConfig()

	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	c.Bind(amigaFwkConfig)

	return c, nil
}

func (c *defaultConfig) reloadConfig() error {
	var k = koanf.New(".")

	// Load environment variables
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "_", ".")
	}), nil); err != nil {
		log.Fatal("error loading env: ", err)
		return fmt.Errorf("error loading env: %w", err)
	}

	// Load configmap.yml
	const configMapFile = "configmap.yml"
	if err := k.Load(file.Provider(configMapFile), yaml.Parser()); err != nil {
		log.Fatal("error loading configmap.yml: ", err)
		return fmt.Errorf("error loading configmap.yml: %w", err)
	}

	// Load secret.yml
	const secretFile = "secret.yml"
	if err := k.Load(file.Provider(secretFile), yaml.Parser()); err != nil {
		log.Fatal("error loading secret.yml: ", err)
		return fmt.Errorf("error loading secret.yml: %w", err)
	}

	// If config now is enabled (amiga.confignow.enabled=true), load configmap from ConfigNow
	if k.Bool("amiga.confignow.enabled") {
		const configNowFile = "confignow.yml"
		f := file.Provider(configNowFile)
		// Load configmap from ConfigNow
		if err := k.Load(f, yaml.Parser()); err != nil {
			log.Fatal("error loading configmap from ConfigNow: ", err)
			return fmt.Errorf("error loading configmap from ConfigNow: %w", err)
		}

		if !c.watcherRegistered {
			f.Watch(func(event any, err error) {
				if err != nil {
					log.Printf("error watching configmap from ConfigNow: %s", err)
					return
				}
				c.onConfigNowReload()
			})
			c.watcherRegistered = true
		}
	}

	c.koanf = k

	return nil
}

// Called when the configmap from ConfigNow is reloaded and the change is detected by the Koanf file watcher
func (c *defaultConfig) onConfigNowReload() {
	log.Println("ConfigNow has been reloaded")

	// TODO the second parameter should be documented as AmigaEvent struct or interface TBD, although the call is made by reflection and here there is no type checking
	c.eventBus.Publish("confignow.refresh", "Event string")

	err := c.reloadConfig()

	if err != nil {
		log.Printf("error reloading config: %s", err)
	}

	err = c.rebind()

	if err != nil {
		log.Printf("error rebinding config: %s", err)
	}
}

func (c *defaultConfig) String(key string) (string, error) {
	if !c.koanf.Exists(key) {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.String(key), nil
}

func (c *defaultConfig) Strings(key string) ([]string, error) {
	if !c.koanf.Exists(key) {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Strings(key), nil
}

func (c *defaultConfig) Int(key string) (int, error) {
	if !c.koanf.Exists(key) {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Int(key), nil
}

func (c *defaultConfig) Ints(key string) ([]int, error) {
	if !c.koanf.Exists(key) {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Ints(key), nil
}

func (c *defaultConfig) Bool(key string) (bool, error) {
	if !c.koanf.Exists(key) {
		return false, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Bool(key), nil
}

func (c *defaultConfig) Bools(key string) ([]bool, error) {
	if !c.koanf.Exists(key) {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Bools(key), nil
}

func (c *defaultConfig) Float64(key string) (float64, error) {
	if !c.koanf.Exists(key) {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Float64(key), nil
}

func (c *defaultConfig) Float64s(key string) ([]float64, error) {
	if !c.koanf.Exists(key) {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return c.koanf.Float64s(key), nil
}

func (c *defaultConfig) Bind(conf AmigaConfig) error {
	if err := unMarshal(conf, c.koanf); err != nil {
		return fmt.Errorf("error binding struct: %w", err)
	}
	c.boundStructs = append(c.boundStructs, &conf)
	return nil
}

func unMarshal(conf AmigaConfig, k *koanf.Koanf) error {
	return k.UnmarshalWithConf(conf.Prefix(), conf, koanf.UnmarshalConf{Tag: "yaml"})
}

func (c *defaultConfig) AmigaFwkConfig() *AmigaFwkConfig {
	return amigaFwkConfig
}

// Rebind all structs that have been bound to the configuration
func (c *defaultConfig) rebind() error {
	totalBoundStructs := len(c.boundStructs)
	if totalBoundStructs > 0 {
		log.Printf("Rebinding %d configuration structs", totalBoundStructs)
		for i := 0; i < totalBoundStructs; i++ {
			if err := unMarshal(*c.boundStructs[i], c.koanf); err != nil {
				return fmt.Errorf("error rebinding struct: %w", err)
			}
		}
	}
	return nil
}

// AmigaFwkConfig is the struct that maps all Amiga Fwk properties (TODO extract from here to Amiga core instance?)
type AmigaFwkConfig struct {
	Amiga struct {
		Common struct {
			Cache struct {
				Redis struct {
					Enabled    bool     `yaml:"enabled"`
					Host       string   `yaml:"host"`
					Port       int      `yaml:"port"`
					Password   string   `yaml:"password"`
					CacheNames []string `yaml:"cache-names"`
				} `yaml:"redis"`
			} `yaml:"cache"`
		} `yaml:"common"`
	} `yaml:"amiga"`
}

// Implementation of AmigaConfig interface for AmigaFwkConfig
func (ac *AmigaFwkConfig) Prefix() string {
	return ""
}
