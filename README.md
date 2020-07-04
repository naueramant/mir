# Mir 

YAML configurable web dashboard viewer written in Go

## Features

- Configurable through YAML
- Multiple tabs cycling
- Basic Auth login
- Injection of custom CSS and JS
- Cron jobs
  - System commands
  - Message flashing

## History

Mir was originally written as an easy and consistent way of configuring web driven dashboards running on a Raspberry PI connected to a TV.

Mir uses YAML to describe it's state and is therefor easy to version control.

## Deployment

A binary for the appropriate platform can be found in [releases](https://github.com/naueramant/mir/releases).

The host system needs Chrome or Chromium to be installed for Mir to work.

Mir will look for a `screen.yaml` file in it's working directory to use as configuration. Mir can also be started with a `--config` flag where a path to a configuration file can be specified.

## Configuration

The Mir configuration file consist of two main sections, *tabs* where the tabs which will be cycled through is specified and *jobs* where cron jobs can be specified.

Configuration examples can be found in [examples](examples). Documentation of configuration options can be found below.

### Tabs

### Jobs