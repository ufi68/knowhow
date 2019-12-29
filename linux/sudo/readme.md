# sudo commands

## change to root 
```
$ sudo -i
```

## Change to root without entry of password
```
$ visudo 
<userID>     ALL=(ALL) NOPASSWD:ALL
```

## How can I make stars appear when I type sudo password?
[askubunu.com](https://askubuntu.com/questions/860766/how-can-i-make-stars-appear-when-i-type-sudo-password)


Locate the line containing ```env_reset ``` and add the parameter ```, pwfeedback ``` behind it. Here is an example:
```
Defaults        env_reset, timestamp_timeout=120, pwfeedback
```

