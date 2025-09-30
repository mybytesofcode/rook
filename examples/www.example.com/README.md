# ~/.ssh/config

```
Match Host "127.0.0.1" exec "knock -d 500 127.0.0.1 2000:tcp 2001:tcp 2002:tcp"
    Hostname 127.0.0.1
    Port 1234
    User user
    RequestTTY yes
```
