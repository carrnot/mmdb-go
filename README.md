# mmdb-go

Enriching MMDB files with your own data using Go

# How to use?

## build
```
go build mmdb-go
```

## example

```bash
curl -L https://raw.githubusercontent.com/carrnot/china-ip-list/release/ip.txt -o ip.txt
./mmdb-go -i ip.txt
```

custom name

```
./mmdb-go -i ip.txt -o custom.mmdb
```

## Credits

* [https://blog.maxmind.com/2020/09/enriching-mmdb-files-with-your-own-data-using-go/](https://blog.maxmind.com/2020/09/enriching-mmdb-files-with-your-own-data-using-go/)