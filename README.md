# davinci-pingcli

## Install

Only Homebrew is supported initially, for other methods of install please [raise an issue](https://github.com/patrickcping/davinci-pingcli/issues/new?title=New%20installation%20method%20required).

### Homebrew

```shell
brew tap patrickcping/davinci-pingcli
brew install davinci-pingcli
```

Test the installation with:

```shell
davinci-pingcli --help
```

## Using `--help`

Help can found by using the `--help` / `-h` parameter for any command.

Examples:

```shell
davinci-pingcli flows list --help
```

## Return Codes

The `davinci-pingcli validate ...` command will return the following return status codes along with a text description of the validation issues:

- `0` - Successful command, no warnings
- `1` - Unsuccessful (errored) command

## Logging

The logging level can be set using the `DAVINCI_PINGCLI_LOG` environment variable.  The possible values that can be set are `DEBUG`, `INFO`, `WARN`, `ERROR` and `NOLEVEL`.  By default, extra log output is disabled.

A log file can be created using the `DAVINCI_PINGCLI_PATH` environment variable.  This redirects the log output to a file of choice.  If not set, the debug output will be printed alongside the command output (stdout).

The following example logs debug output to the file called `davinci-pingcli.log`:
```shell
DAVINCI_PINGCLI_LOG=DEBUG DAVINCI_PINGCLI_PATH=`pwd`/davinci-pingcli.log davinci-pingcli flows list
```

## Configuration File

The parameters described above can be configured in a static configuration file, expected to be named `.davinci-pingcli.yaml` in the same directory that the CLI tool is run.  The following describe the properties that can be set, and an example can be found at [./davinci-pingcli.example](./blob/main/.davinci-pingcli.example)

#### General Properties

The following are configuration file settings for the commands.

| Config File Property | Environment Variable | Type          | Equivalent Parameter        | Purpose                                                               |
|----------------------|---------------|---------------|-----------------------------|-----------------------------------------------------------------------|
| `username`     | `PINGCLI_DAVINCI_USERNAME` | string  | `--username` / `-u` | The admin username used to connect to DaVinci.                             |
| `password`     | `PINGCLI_DAVINCI_PASSWORD` | string  | `--password` / `-p` | The admin password used to connect to DaVinci.                             |
| `adminEnvironmentId`     | `PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID` | UUID  | `--admin-environment-id` / `-e` | The PingOne environment ID that contains the admin user.                             |
| `environmentId`     | `PINGCLI_DAVINCI_ENVIRONMENT_ID` | UUID  | `--environment-id` / `-t` | The PingOne environment ID to control configuration for.                             |
| `region`     | `PINGCLI_DAVINCI_REGION` | string  | `--region` / `-r` | The region where the PingOne environment is located.  Options are `AsiaPacific`, `Canada`, `Europe` and `NorthAmerica`.                             |
