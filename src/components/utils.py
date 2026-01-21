import os
import json
import logging

logger = logging.getLogger(__name__)

def get_redis_config(filename):
    """Get Redis configuration from a file."""
    if not os.path.exists(filename):
        raise FileNotFoundError(f"File {filename} not found")

    with open(filename, 'r') as f:
        config = json.load(f)

    return config

def get_environment_variables():
    """Get environment variables."""
    env_variables = {}
    for key, value in os.environ.items():
        env_variables[key] = value
    return env_variables

def get_cache_config():
    """Get cache configuration from environment variables and Redis configuration."""
    config = get_environment_variables()
    if 'REDIS_HOST' not in config:
        raise ValueError("REDIS_HOST environment variable not set")

    if 'REDIS_PORT' not in config:
        raise ValueError("REDIS_PORT environment variable not set")

    redis_config = get_redis_config('/etc/redis/config.json')
    redis_config['host'] = config['REDIS_HOST']
    redis_config['port'] = config['REDIS_PORT']

    return redis_config