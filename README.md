Simple metrics sink to push data into (currently) InfluxDB.

## Description
It is meant as a simple fire-and-forget library, configured from environment variables. If none are set, it will silently drop all data. Please be aware that you are expected to use the go keyword when pushing metrics unless a wait of a few seconds is no problem.

## Usage
```golang
go metrics.WithTag(
    "customer",
    "acme dynamite inc.",
).Publish(
    "explosion",
    metrics.Float32("megajoule", 12.3),
    metrics.Int("coffee", 2),
)
```

The library adds some tags automatically:
  - the hostname the program is running on as "host"
  - the application name argv[0] as "application"

## InfluxDB
Currently InfluxDB is the only supported backend. Patches welcome!

In influx each measurements is split into data series identified by the name and all tags. By default there can be up to 1.000.000 series in one database.

### Using the InfluxDB backend
To enable the InfluxDB backend include it in your program
```golang
import "github.com/jpicht/metrics/sink/influx"
```

And provide the configuration via the environment
```bash
export INFLUX_SERVER="http://influx.your.local.domain:8086/"
export INFLUX_DATABASE="metrics"
```

The database needs to exist, otherwise no data will be recorded. The library will silently swallow all error messages. Please make sure to also set up a retention policy, so data will not be kept indefinitely.
