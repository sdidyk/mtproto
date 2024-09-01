# telegram
MTProto implementation in Golang with example tool.

# Example tool

Example tool saves the authkey and other data in ~/.telegram_go. If you delete/lost this file, you will need to auth again.

## install

```
$ go install github.com/sdidyk/mtproto/example/telegram@latest
```

## phone authorization

Authorizes phone by code (sent by sms or thru telegram app).

```
$ telegram auth 79998886655
Enter code: 25704
Signed in: id 104842610 name <Sergei Didyk>
```

## get contact list

Shows user's contact list.

```
$ telegram list
        id      mutual    name                              username                               access_hash
     40352        true    Николай ***                       n****                             94c173dd********
     88928        true    Сергей ***                        s****                             c72a5847********
    109551        true    Мария ***                         m****                             ff92275a********
```

## send message to contact

Sends a message to contact. Destination id should be from contact list or user himself.

```
$ telegram msg 109551 ff92275a******** 'Hack the planet!'
```

## Library

*documentation not ready yet*

### MTProto documentation
* https://core.telegram.org/mtproto
* https://core.telegram.org/api

## License

MIT License
