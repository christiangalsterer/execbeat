[![Build Status](https://travis-ci.org/christiangalsterer/execbeat.svg?branch=i37)](https://travis-ci.org/christiangalsterer/execbeat)
[![codecov.io](http://codecov.io/github/christiangalsterer/execbeat/coverage.svg?branch=master)](http://codecov.io/github/christiangalsterer/execbeat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/christiangalsterer/execbeat)](https://goreportcard.com/report/github.com/christiangalsterer/execbeat)
[![license](https://img.shields.io/github/license/christiangalsterer/execbeat.svg)](https://github.com/christiangalsterer/execbeat)
[![Github All Releases](https://img.shields.io/github/downloads/christiangalsterer/execbeat/total.svg)](https://github.com/christiangalsterer/execbeat)

![Elastic Beats 7.10.1](https://img.shields.io/badge/Elastic%20Beats-v7.10.1-blue.svg)
![Golang 1.15](https://img.shields.io/badge/Golang-v1.15-blue.svg)

# Execbeat

Execbeat is the [Beat](https://www.elastic.co/products/beats) used to execute any command.
Multiple commands can be configured which are executed in a regular interval and the standard output and standard error is shipped to the configured output channel.

Execbeat is inspired by the Logstash [exec](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-exec.html) input filter but doesn't require that the endpoint is reachable by Logstash as Execbeat pushes the data to Logstash or Elasticsearch.
This is often necessary in security restricted network setups, where Logstash is not able to reach all servers. Instead the server to be monitored itself has Execbeat installed and can send the data or a collector server has Execbeat installed which is deployed in the secured network environment and can reach all servers to be monitored.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/christiangalsterer/execbeat`

## Getting Started with Execbeat

### Requirements

- [Golang](https://golang.org/dl/) 1.15

### Init Project

To get running with Execbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Execbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/christiangalsterer/execbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Execbeat run the command below. This will generate a binary
in the same directory with the name execbeat.

```
make
```

### Run

To run Execbeat with debugging output enabled, run:

```
./execbeat -c execbeat.yml -e -d "*"
```

### Test

To test Execbeat, run the following command:

```
make testsuite
```

alternatively:

```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```

### Cleanup

To clean Execbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

### Clone

To clone Execbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/christiangalsterer/execbeat
git clone https://github.com/christiangalsterer/execbeat ${GOPATH}/src/github.com/christiangalsterer/execbeat
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

# Releases

## 3.3.0 (2017-10-06) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.3.0)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.3.0/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/3.3.0)

Feature and Bugfix release containing the following changes:

- Update to beats v5.6.2

## 3.2.0 (2017-06-05) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.2.0)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.2.0/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/3.2.0)

Feature and bugfix release containing the following changes:

- Fix: [Use exit code 127 when command to execute is not found](https://github.com/christiangalsterer/execbeat/issues/15)

## 3.1.1 (2017-02-24) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.1.1)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.1.1/total.svg)](https://g0thub.com/christiangalsterer/execneat/releases/tag/3.1.1)

Bugfix release containing the following changes:

- [Set correct version in package names and package metadata](https://github.com/christiangalsterer/execbeat/issues/10)

## 3.1.0 (2017-02-23) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.1.0)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.1.0/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/3.1.0)

Feature and bugfix release containing the following changes:

- The exit code of the command executed is now exported in field `exitCode`.
- Fix: Examples were not fully updated with configuration changes introduced in 3.0.0.

## 3.0.1 (2017-02-21) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.1)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.0.1/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/3.0.1)

Bugfix release containing the following changes:

- [Multiple arguments are not properly passed](https://github.com/christiangalsterer/execbeat/issues/7)

## 3.0.0 (2017-02-19) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.0)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.0.0/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/3.0.0)

Feature and bugfix release containing the following **breaking** changes:

- Renamed configuration parameter `execs` to `commands`. Please update your configuration accordingly.
- Renamed configuration parameter `cron` to `schedule`. Please update your configuration accordingly.
- Update to beats v5.2.1
- Fix: [Default schedule not working](https://github.com/christiangalsterer/execbeat/issues/6)

## 2.2.0 (2017-02-04) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/2.2.0)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/2.2.0/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/2.2.0)

Feature release containing the following changes:

- Update to beats v5.2.0

## 2.1.1 (2017-01-14) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/2.1.1)

[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/2.1.1/total.svg)](https://github.com/christiangalsterer/execneat/releases/tag/2.1.1)

Starting with this release pre-compiled binaries for different operating systems are available under the respective tag in the github project.

Bugfix release containing the following changes:

- Move files into correct place to allow correct bulding with `make package`
- Move files into correct place to allow correct bulding with `make update`
- Cleanup of documentation
- Update to beats v5.1.2
- Update to Go 1.7.4

## 2.1.0 (2016-12-23)

Feature release containing the following changes:

- Update to beats v5.1.1

## 2.0.0 (2016-11-26)

Feature release containing the following changes:

- Update to beats v5.0.1

Please note that this release contains the following breaking changes introduced by beats 5.0.X, see also [Beats Changelog](https://github.com/elastic/beats/blob/v5.0.0-beta1/CHANGELOG.asciidoc)

- SSL Configuration
  - rename tls configurations section to ssl
  - rename certificate_key configuration to key.
  - replace tls.insecure with ssl.verification_mode setting.
  - replace tls.min/max_version with ssl.supported_protocols setting requiring full protocol name

## 1.1.0 (2016-07-19)

Feature release containing the following changes:

- Update to Go 1.6
- Update to libbeat 1.2.3
- Use [Glide](https://github.com/Masterminds/glide) for dependency management

## 1.0.1 (2016-02-15)

Bugfix release containing the following changes:

- Fix: [Hanging during shutdown](https://github.com/christiangalsterer/execbeat/issues/2)

## 1.0.0 (2015-12-26)

- Initial release

# Configuration

## Configuration Options

See [here](docs/configuration.asciidoc) for more information.

## Exported Document Types

There is exactly one document type exported:

- `type: execbeat` command execution information, e.g. standard output and standard error. The type can be changed by setting the document_type attribute.

## Exported Fields

See [here](docs/fields.asciidoc) for a detailed description of all exported fields.

### execbeat type

<pre>
{
  "_index": "execbeat-2015.12.26",
  "_type": "execbeat",
  "_source": {
    "@timestamp": "2015-12-26T02:18:53.001Z",
    "beat": {
      "hostname": "mbp.box",
      "name": "mbp.box"
    },
    "count": 1,
    "fields": {
      "host": "test"
    },
    "exec": {
      "command": "echo",
      "exitCode": 0,
      "stdout": "Hello World\n"
    },
    "fields": {
      "host": "test2"
    },
    "type": "execbeat"
    },
  "sort": [
    1449314173
  ]
}
</pre>

## Elasticsearch Template

To apply the Execbeat template:

    curl -XPUT 'http://localhost:9200/_template/execbeat' -d@etc/execbeat.template.json

# Contribution

All sorts of contributions are welcome. Please create a pull request and/or issue.
