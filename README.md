[![Build Status](https://travis-ci.org/christiangalsterer/execbeat.svg?branch=master)](https://travis-ci.org/christiangalsterer/execbeat)
[![codecov.io](http://codecov.io/github/christiangalsterer/execbeat/coverage.svg?branch=master)](http://codecov.io/github/christiangalsterer/execbeat?branch=master)

# Overview

Execbeat is the [Beat](https://www.elastic.co/products/beats) used to execute any command.
Multiple commands can be configured which are executed in a regular interval and the standard output and standard error is shipped to the configured output channel.

Execbeat is inspired by the Logstash [exec](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-exec.html) input filter but doesn't require that the endpoint is reachable by Logstash as Execbeat pushes the data to Logstash or Elasticsearch.
This is often necessary in security restricted network setups, where Logstash is not able to reach all servers. Instead the server to be monitored itself has Execbeat installed and can send the data or a collector server has Execbeat installed which is deployed in the secured network environment and can reach all servers to be monitored.

# Configuration

## Configuration Options

## Exported Document Types

There is exactly one document type exported:

- `type: execbeat` http request and response information

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

# Build, Test, Run

```
# Build
GOPATH=<your go path> make execbeat

# Test
GOPATH=<your go path> make test

# Run
./execbeat -c /etc/execbeat/execbeat.yml
```
# Contribution
All sorts of contributions are welcome. Please create a pull request and/or issue.
