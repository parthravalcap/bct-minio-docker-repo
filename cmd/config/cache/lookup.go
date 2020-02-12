/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import (
	"errors"
	"strconv"

	"github.com/minio/minio/cmd/config"
	"github.com/minio/minio/pkg/env"
)

// Cache ENVs
const (
	Drives  = "drives"
	Exclude = "exclude"
	Expiry  = "expiry"
	MaxUse  = "maxuse"
	Quota   = "quota"
	After   = "after"

	EnvCacheDrives  = "MINIO_CACHE_DRIVES"
	EnvCacheExclude = "MINIO_CACHE_EXCLUDE"
	EnvCacheExpiry  = "MINIO_CACHE_EXPIRY"
	EnvCacheMaxUse  = "MINIO_CACHE_MAXUSE"
	EnvCacheQuota   = "MINIO_CACHE_QUOTA"
	EnvCacheAfter   = "MINIO_CACHE_AFTER"

	EnvCacheEncryptionMasterKey = "MINIO_CACHE_ENCRYPTION_MASTER_KEY"

	DefaultExpiry = "90"
	DefaultQuota  = "80"
	DefaultAfter  = "0"
)

// DefaultKVS - default KV settings for caching.
var (
	DefaultKVS = config.KVS{
		config.KV{
			Key:   Drives,
			Value: "",
		},
		config.KV{
			Key:   Exclude,
			Value: "",
		},
		config.KV{
			Key:   Expiry,
			Value: DefaultExpiry,
		},
		config.KV{
			Key:   Quota,
			Value: DefaultQuota,
		},
		config.KV{
			Key:   After,
			Value: DefaultAfter,
		},
	}
)

const (
	cacheDelimiter = ","
)

// Enabled returns if cache is enabled.
func Enabled(kvs config.KVS) bool {
	drives := kvs.Get(Drives)
	return drives != ""
}

// LookupConfig - extracts cache configuration provided by environment
// variables and merge them with provided CacheConfiguration.
func LookupConfig(kvs config.KVS) (Config, error) {
	cfg := Config{}

	if err := config.CheckValidKeys(config.CacheSubSys, kvs, DefaultKVS); err != nil {
		return cfg, err
	}

	drives := env.Get(EnvCacheDrives, kvs.Get(Drives))
	if len(drives) == 0 {
		return cfg, nil
	}

	var err error
	cfg.Drives, err = parseCacheDrives(drives)
	if err != nil {
		return cfg, err
	}

	cfg.Enabled = true
	if excludes := env.Get(EnvCacheExclude, kvs.Get(Exclude)); excludes != "" {
		cfg.Exclude, err = parseCacheExcludes(excludes)
		if err != nil {
			return cfg, err
		}
	}

	if expiryStr := env.Get(EnvCacheExpiry, kvs.Get(Expiry)); expiryStr != "" {
		cfg.Expiry, err = strconv.Atoi(expiryStr)
		if err != nil {
			return cfg, config.ErrInvalidCacheExpiryValue(err)
		}
	}

	if maxUseStr := env.Get(EnvCacheMaxUse, kvs.Get(MaxUse)); maxUseStr != "" {
		cfg.MaxUse, err = strconv.Atoi(maxUseStr)
		if err != nil {
			return cfg, config.ErrInvalidCacheQuota(err)
		}
		// maxUse should be a valid percentage.
		if cfg.MaxUse < 0 || cfg.MaxUse > 100 {
			err := errors.New("config max use value should not be null or negative")
			return cfg, config.ErrInvalidCacheQuota(err)
		}
		cfg.Quota = cfg.MaxUse
	} else if quotaStr := env.Get(EnvCacheQuota, kvs.Get(Quota)); quotaStr != "" {
		cfg.Quota, err = strconv.Atoi(quotaStr)
		if err != nil {
			return cfg, config.ErrInvalidCacheQuota(err)
		}
		// quota should be a valid percentage.
		if cfg.Quota < 0 || cfg.Quota > 100 {
			err := errors.New("config quota value should not be null or negative")
			return cfg, config.ErrInvalidCacheQuota(err)
		}
		cfg.MaxUse = cfg.Quota
	}

	if afterStr := env.Get(EnvCacheAfter, kvs.Get(After)); afterStr != "" {
		cfg.After, err = strconv.Atoi(afterStr)
		if err != nil {
			return cfg, config.ErrInvalidCacheAfter(err)
		}
		// after should be a valid value >= 0.
		if cfg.After < 0 {
			err := errors.New("cache after value cannot be less than 0")
			return cfg, config.ErrInvalidCacheAfter(err)
		}
	}

	return cfg, nil
}
