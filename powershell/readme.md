
## Power Shell

### Windows Power Shell 

### PowerShell Core
Windos, Linux and macOS

## Show Installed Version
```
PS> $PSVersionTable

Name                           Value
----                           -----
PSVersion                      5.1.18362.145
PSEdition                      Desktop
PSCompatibleVersions           {1.0, 2.0, 3.0, 4.0...}
BuildVersion                   10.0.18362.145
CLRVersion                     4.0.30319.42000
WSManStackVersion              3.0
PSRemotingProtocolVersion      2.3
SerializationVersion           1.1.0.1
```

## Show Installed Version.Version 
```
PS> $PSVersionTable.$PSVersion

Major  Minor  Build  Revision
-----  -----  -----  --------
5      1      18362  145
```

## Get installed Provider
```
PS> Get-PSProvider
Name                 Capabilities                                      Drives
----                 ------------                                      ------
Registry             ShouldProcess, Transactions                       {HKLM, HKCU}
Alias                ShouldProcess                                     {Alias}
Environment          ShouldProcess                                     {Env}
FileSystem           Filter, ShouldProcess, Credentials                {C, H, E, F...}
Function             ShouldProcess                                     {Function}
Variable             ShouldProcess                                     {Variable}
```
## Get availabe Commands
```
PS>  Get-command
```

# Powershell on Linux

```
$ pwsh -v 
```