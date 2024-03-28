# How to version your golang app 

## Tag 

```
git tag v1.0.0 # tag current commit 
git tag -n # show tags and respective commits
```

## ldflags

```
go build -ldflags="-X 'package_path.variable_name=new_value'"
# example 
go build -o app -ldflags="-X 'main.version=my_version'"
```