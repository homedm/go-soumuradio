# go-soumuradio

soumuradio is a Go client library for accessing the radio station etc.
information search API of Ministry of Internal Affairs and Communications in
Japan

## Install

```sh
go get github.com/tomato3713/go-soumuradio
```

## Usage

```sample01.go
import github.com/tomato3713/go-soumuradio/soumu
```

```
cli, err := soumu.NewClient("")
if err != nil {
    os.Exit(1)
}

opts := soumu.NewNumOptions(soumu.License, soumu.Amateur)
result, err := cli.GetNum(nil, opts)
if err != nil {
    os.Exit(1)
}

fmt.Printf("%+v", *result)
// -> {Musen:{Count:404542} MusenInformation:{TotalCount:404542 LastUpdateDate:2019-11-03}}
```

## Note

このサービスは、総務省 電波利用ホームページのWeb-API 機能を利用して取得した情報
をもとに作成しているが、サービスの内容は総務省によって保証されたものではない

## Develope Note

- logger: Use logger interface defined in logger.go
