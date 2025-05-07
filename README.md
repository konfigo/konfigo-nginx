# konfigo-nginx

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

NGINX wrapper that integrates configurations from Konfigo

## Instructions

- Consider any variable `VAR` that is stored in the static distributable to be hosted in the nginx server as `__VAR__` will be actively parsed as the `VAR` value according to path in the configuration

## Environment Variables

| Environment Variable   | Description                                                                        | Example                     |
| ---------------------- | ---------------------------------------------------------------------------------- | --------------------------- |
| `KONFIGO_API_ENDPOINT` | API endpoint to fetch configuration from                                           | `http://localhost:3000/api` |
| `KONFIGO_API_KEY`      | API key to use when fetching configuration                                         | `super_secure_api_key`      |
| `KONFIGO_PATH`         | Path to fetch configuration from                                                   | `app/dev/v1`                |
| `KONFIGO_INTERVAL`     | Interval in seconds to fetch the configuration. Defaults to 10 seconds if not set. | 10                          |
