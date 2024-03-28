# How to version your golang app 

## Tag 

```
git tag v1.0.0 # tag current commit 
git tag -n # show tags and respective commits
git describe --tags # describe lastest tag in the branch 
# if current commit is commit of a tag then 
v1.0.1
# if current commit is not commit of a tag
v1.0.1-2-g21ae9b6
# v1.0.1 - tag
# 2 - how far away the current commit from tagged commit
# g21ae9b6 - g prefix for "git", 21ae9b6 is short hash  
```

## ldflags

```
go build -ldflags="-X 'package_path.variable_name=new_value'"
# example 
go build -o app -ldflags="-X 'main.version=my_version'"
```

## how to get version? 
```
# in root of this repo
./compile.sh 
./app 
./app version 
```