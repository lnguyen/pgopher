# pgopher

Runs single postgres command in static binary.

Sometimes you just need to run postgres command and don't want to install postgres on machine.

```
pgopher --host "localhost" --port 5432 -u lnguyen -p lnguyen --database lnguyen-q "create table foo ()"
# Or
pgopher --uri "postgres://lnguyen:lnguyen@localhost:5432/lnguyen" -q "create table foo ()"
```

## Installation

```
go get github.com/longnguyen11288/pgopher
```

## TODO

Add support for running select query and printing them out
