# myip

This is a stupid Go program to simply fetch my IP.

## caveats

It has the following characteristics idiosyncratic to my setup:

1. It excludes my local Docker address (any address whose subnet includes 172.17.0.1)

2. With the `-x` flag, it uses a script on my website to fetch the "external" IP.  (e.g. the address of my DSL when at home)

## usage

```
go run myip.go
```

## installation

```
## build it
go build

## copy it to your $PATH
cp myip ~/bin/
```

## flags

```
-6   will prefer IPv6 (default prefers IPv4)
-x   will fetch the "external" IP
```
