# telegram
MTProto implementation in Golang with example tool.

# Example tool

Example tool saves the authkey and other data in ~/.telegram_go. If you delete/lost this file, you will need to auth again.

## install

```
$ go get -v -u github.com/sdidyk/mtproto/example/telegram
```

## phone authorization

Authorizes phone by code (sent by sms or thru telegram app).

```
$ telegram auth 79998886655
Enter code: 25704
Signed in: id 104842610 name <Sergey Didyk>
```

## get contact list

Shows user's contact list.

```
$ telegram list
        id        mutual    name                              username
    132597         false    Алексей Г*******                  O******
    326007         false    Татьяна К*******
    344375          true    Андрей П*********                 r******
    348798          true    Руслан Ч******
    473977         false    Тимур Д******
```

## send message to contact

Sends a message to contact. Destination id should be from contact list or user himself.

```
$ telegram msg 104842610 'Hack the planet!'
```

## Library

*documentation not ready yet*

### MTProto documentation
* https://core.telegram.org/mtproto
* https://core.telegram.org/api

## License

MIT License
