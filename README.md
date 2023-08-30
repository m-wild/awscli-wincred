# awscli-wincred

### Build

```
PS> go build
```

### Usage

```
PS> ./awscli-wincred.exe --help
Usage of awscli-wincred.exe:
  -get
        retrieve credentials from the store
  -profile string
        aws cli profile name
  -set
        save or update credentials in the store
```

### Get/set credentials

```
PS> ./awscli-wincred.exe --profile=personal --set
AccessKeyID: example
SecretAccessKey: example

PS> ./awscli-wincred.exe --profile=personal --get
{
  "Version": 1,
  "AccessKeyId": "example",
  "SecretAccessKey": "example"
}
```

### Configure AWS CLI

e.g. $PROFILE/.aws/credentials
```
[personal]
credential_process = "C:\Path\To\awscli-wincred.exe" --profile=personal --get
```


