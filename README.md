# Mir

YAML configurable web dashboard viewer written in Go

## Features

- Fully declarable in YAML
- Multiple tabs cycling
- Basic Auth login
- Injection of custom CSS and JS
- Cron jobs
  - System commands
  - Message flashing
- Detect failed page loads and retry

## About

Mir is an easy and consistent way of configuring web driven dashboards originally created to run on a Raspberry PI connected to a TV and display dashboards such a grafana etc.

Mir is fully declarable in YAML and is therefor easy to version control and copy to new screens.

## Deployment

Binary releases can be found in [releases](https://github.com/naueramant/mir/releases).

The system where Mir will be running needs Chrome or Chromium to be installed.

Mir will look for a `screen.yaml` file in it's working directory to use as configuration. Mir can also be started with a `--config` flag where a configuration file can be specified ex.

```sh
mir --config foo/bar.yaml
```

## Configuration

The Mir configuration file consist of two main sections, _tabs_ where the tabs which will be cycled through is specified and _jobs_ where cron jobs can be specified. A simple example could one tab with xkcd comics and rebooting the system every night at 03:00, which would look like this:

```yaml
syntax: v1

tabs:
  - url: https://xkcd.com/

jobs:
  - type: command
    when: "0 0 3 * *"
    options:
      command: sudo
      args:
        - reboot
```

More configuration examples can be found in [examples](examples). Full documentation of the configuration can be found below:

| name   | type   | usage                                                                     |
| ------ | ------ | ------------------------------------------------------------------------- |
| syntax | string | The syntax version of the configuration, only valid version is v1 for now |
| tabs   | Tab[]  | An array of _Tab_ configurations                                          |
| jobs   | Job[]  | An array of _Job_ configurations                                          |

**Tab**:

| name     | type   | usage                                                                                                                       |
| -------- | ------ | --------------------------------------------------------------------------------------------------------------------------- |
| url      | string | The URL which the tab will load                                                                                             |
| duration | number | _(optional)_ If more than one tab is specified the duration will be the number of seconds before the next tab will be shown |
| reload   | bool   | _(optional)_ Reload the tab before it is switched to                                                                        |
| css      | string | _(optional)_ Path to a css file which should be injected into the tab                                                       |
| js       | string | _(optional)_ Path to a js file which should be injected into the tab                                                        |
| auth     | Auth   | _(optional)_ Basic auth login options                                                                                       |

**Auth**:

| name     | type   | usage               |
| -------- | ------ | ------------------- |
| username | string | Basic Auth username |
| password | string | Basic Auth password |

### Jobs

The _jobs_ field is an array of _job_ configurations:

**Job**:

| name    | type    | usage                                                                           |
| ------- | ------- | ------------------------------------------------------------------------------- |
| type    | string  | Job type identifier, see Options for available types                            |
| when    | string  | cron job expression in format minute, hour, day of month, month and day of week |
| options | Options | A option object for the specific job type                                       |

**Options**:

**Type**: _command_

| name    | type     | usage                      |
| ------- | -------- | -------------------------- |
| command | string   | Command to run             |
| args    | string[] | Array of command arguments |

**Type**: _tab_

| name     | type     | usage                                                |
| -------- | -------- | ---------------------------------------------------- |
| url      | string   | The URL which the tab will load                      |
| duration | number   | Duration in seconds of which the tab will be visible |

**Type**: _message_

| name            | type   | usage                                                                         |
| --------------- | ------ | ----------------------------------------------------------------------------- |
| duration        | number | Duration in seconds of which the message will be visible                      |
| message         | string | Text message to display                                                       |
| fontSize        | number | Message font size in pixel                                                    |
| textColor       | string | Text color                                                                    |
| backgroundColor | string | Background color                                                              |
| blink           | bool   | Switch text color and background color every second to make the message blink |
