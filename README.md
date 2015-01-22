# check_growthforecast_value

Monitoring [GrowthForecast](http://kazeburo.github.io/GrowthForecast/) current value from [Nagios](http://www.nagios.org/)(plugin)

## Install

You can get binary from [releases](https://github.com/hirocaster/check_growthforecast_value/releases).

## Usage

greater than value(option), `Warning` or `Critical` case ...

```
$ check_growthforecast_value  -u "http://growthforcast.domain.host" -i "service/section/name" -w 70 -c 90
```

less than value(option),  `Warning` or `Critical` case ...

```
$ check_growthforecast_value  -u "http://growthforcast.domain.host" -i "service/section/name" -direction "downward" -w 30 -c 10
```

Please, check option details

```
$ check_growthforecast_value --help
```
