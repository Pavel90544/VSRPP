# Redis & Memory Cache

[![CI](https://github.com/Pavel90544/VSRPP/actions/workflows/cache-ci.yml/badge.svg)](https://github.com/Pavel90544/VSRPP/actions/workflows/cache-ci.yml)

## Cache Implementations

- **MemoryCache** - In-memory cache with TTL
- **RedisCache** - Redis-based cache

## Quick Start

```bash
cd lab5
go test -v ./internal/pkg/cache/...