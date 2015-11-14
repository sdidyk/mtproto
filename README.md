# telegram
MTProto implementation in Golang with example tool.

# Example tool

Example tool saves the authkey and other data in ~/.telegram_go. If you delete/lost this file, you will need to auth again.

## install

```
$ go get -v -u github.com/hugozhu/mtproto/example/telegram
```

## proxy setting
Socks5 proxy is supported by env variable socks5_proxy
```
$ export socks5_proxy=192.168.1.4:1080
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

## get dialogs

Shows user's dialogs.

```
$ telegram dialogs
        id          type    top_message    unread_count	title               
  20071829	    User	9425      	1    	Raspberry Pi(hugozhu2)
  15626832	    Chat	8896      	0    	树莓派通知               
  69443043	    User	8872      	0    	Hugo Zhu(hugozhu)   
    777000	    User	8871      	0    	Telegram ()         
    333000	    User	8624      	0    	Telegram () 
```


## send message to contact

Sends a message to contact. Destination id should be from contact list or user himself.

```
$ telegram msg 104842610 'Hack the planet!'
```

## send photo to contact or dialog(@dialog_id)

```
$ telegram sendmedia "@20071829" ~/Pictures/IMG_1851.JPG
```

## Library

*documentation not ready yet*

### MTProto documentation
* https://core.telegram.org/mtproto
* https://core.telegram.org/api

## License

MIT License
