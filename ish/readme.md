# iSH - The Linux shell for iOS

https://ish.app/
https://github.com/ish-app/ish

c't 2020 Heft 25 S.166ff

## Install apk 
https://github.com/ish-app/ish/wiki/Installing-apk

`wget -qO- http://dl-cdn.alpinelinux.org/alpine/v3.12/main/x86/apk-tools-static-2.10.5-r1.apk | tar -xz sbin/apk.static && ./sbin/apk.static add apk-tools && rm sbin/apk.static && rmdir sbin 2> /dev/null`

`apk add curl openssh nano`

## Problems with go 
21.02.2023

https://github.com/ish-app/ish/issues/1230#issuecomment-1246190772

> Good day, lads! I've recently successfully compile a simple "Golang hello world" inside the iSH, with the version of go is 1.18.6, for sure not build-from source, just use the command

`apk add go --repository=http://dl-cdn.alpinelinux.org/alpine/latest-stable/community`

> You may use the --repository option to choose different version of go, without the need to build from source.

