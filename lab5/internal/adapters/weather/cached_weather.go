package weather

import (
	"fmt"
	"time"

	"github.com/Pavel90544/VSRPP/lab5/internal/domain/models"
	"github.com/Pavel90544/VSRPP/lab5/internal/pkg/cache"
)

type CachedWeatherInfo struct {
	weather WeatherInfo
	cache   cache.Cache
	ttl     time.Duration
	logger  Logger
}

func NewCachedWeatherInfo(weather WeatherInfo, cache cache.Cache, ttl time.Duration, logger Logger) *CachedWeatherInfo {
	return &CachedWeatherInfo{
		weather: weather,
		cache:   cache,
		ttl:     ttl,
		logger:  logger,
	}
}

func (c *CachedWeatherInfo) GetTemperature(lat, long float64) models.TempInfo {
	cacheKey := fmt.Sprintf("weather:%.4f:%.4f", lat, long)

	c.logger.Debug(fmt.Sprintf("Checking cache for key: %s", cacheKey))

	var cachedTemp models.TempInfo
	if err := c.cache.Get(cacheKey, &cachedTemp); err == nil {
		c.logger.Info(fmt.Sprintf("Cache hit! Temperature: %.2f°C", cachedTemp.Temp))
		return cachedTemp
	}

	c.logger.Debug("Cache miss, fetching from API...")

	temp := c.weather.GetTemperature(lat, long)

	if err := c.cache.Set(cacheKey, temp, c.ttl); err != nil {
		c.logger.Error("Failed to save to cache", err)
	}

	return temp
}
