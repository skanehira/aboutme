# aboutme
aboutme is a profile cli tool.
You can use it by simply writing information in $HOME/.aboutme.yaml.

## Installation
```
# install command
go get -u github.com/skanehira/aboutme
cp aboutme.yaml ~/.aboutme.yaml

# edit your profile
vim ~/.aboutme.yaml
```

## Usage
```
# all of profile
$ aboutme -all

# about sns
$ about -sns

# other options
$ aboutme -h
Usage of aboutme:
  -all
        all of profile
  -db
        about db
  -fw
        about framework
  -job
        about job history
  -lang
        about programming language
  -os
        about os
  -sns
        about sns
  -stat
        about status
  -tool
        about tool
```

## Contribute
If you want to add a new profile item please send me pull requests.
