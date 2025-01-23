package usecase

import (
	"context"
	"log"

	amiga "github.com/JoseLora/fiberapp/internal/amiga/config"
	"github.com/JoseLora/fiberapp/internal/domain/entity"
	"github.com/JoseLora/fiberapp/internal/domain/repository"
	"github.com/JoseLora/fiberapp/internal/domain/usecase"
	"github.com/asaskevich/EventBus"
)

type RedisConf struct {
	Enabled    bool     `yaml:"enabled"`
	Host       string   `yaml:"host"`
	Port       int      `yaml:"port"`
	Password   string   `yaml:"password"`
	CacheNames []string `yaml:"cache-names"`
}

func (r *RedisConf) Prefix() string {
	return "amiga.common.cache.redis"
}

// ProductFinderAll is an implementation of the ProductFinderAll use case.
// It uses a repository to find all products.
type ProductFinderAll struct {
	repository  repository.Product
	amigaConfig amiga.Config
	redisConfig *RedisConf
}

// NewProductFinderAll creates a new instance of ProductFinderAllImpl.
// It takes a repository as a parameter and returns a ProductFinderAll use case.
func NewProductFinderAll(repository repository.Product, amigaConfig amiga.Config, eventBus EventBus.Bus) usecase.ProductFinderAll {
	redisConf := &RedisConf{}
	amigaConfig.Bind(redisConf)

	eventBus.Subscribe("confignow.refresh", func(description string) {
		log.Printf("Received event: confignow.refresh %s", description)
	})

	return &ProductFinderAll{
		repository:  repository,
		amigaConfig: amigaConfig,
		redisConfig: redisConf,
	}
}

// FindAll retrieves all products from the repository.
// It takes a context as a parameter and returns a slice of products and an error.
func (p *ProductFinderAll) FindAll(ctx context.Context) ([]entity.Product, error) {
	conf := p.amigaConfig.AmigaConfigProps()
	log.Println("---- AmigaConfig bound struct ----")
	log.Println(conf.Amiga.Common.Cache.Redis.Enabled)
	log.Println(conf.Amiga.Common.Cache.Redis.Host)
	log.Println(conf.Amiga.Common.Cache.Redis.Port)
	log.Println(conf.Amiga.Common.Cache.Redis.Password)
	log.Println(conf.Amiga.Common.Cache.Redis.CacheNames)

	log.Println("----- AmigaConfig methods -----")
	log.Println(p.amigaConfig.Bool("amiga.common.cache.redis.enabled"))
	log.Println(p.amigaConfig.String("amiga.common.cache.redis.host"))
	log.Println(p.amigaConfig.Int("amiga.common.cache.redis.port"))
	log.Println(p.amigaConfig.String("amiga.common.cache.redis.password"))
	cachenames, _ := p.amigaConfig.Strings("amiga.common.cache.redis.cache-names")
	log.Println(cachenames[0])
	log.Println(cachenames[1])
	log.Println(cachenames[2])

	log.Println("------ Custom bound Redis struct -----")
	log.Println(p.redisConfig.Enabled)
	log.Println(p.redisConfig.Host)
	log.Println(p.redisConfig.Port)
	log.Println(p.redisConfig.Password)
	log.Println(p.redisConfig.CacheNames)

	return p.repository.FindAll(ctx)
}

// ProductFinderByID is an implementation of the ProductFinderByID use case.
// It uses a repository to find a product by its ID.
type ProductFinderByID struct {
	repository repository.Product
}

// NewProductFinderByID creates a new instance of ProductFinderByIDImpl.
// It takes a repository as a parameter and returns a ProductFinderByID use case.
func NewProductFinderByID(repository repository.Product) usecase.ProductFinderByID {
	return &ProductFinderByID{
		repository: repository,
	}
}

// FindByID retrieves a product by its ID from the repository.
// It takes a context and an ID as parameters and returns a product and an error.
func (p *ProductFinderByID) FindByID(ctx context.Context, id string) (entity.Product, error) {
	return p.repository.FindByID(ctx, id)
}
