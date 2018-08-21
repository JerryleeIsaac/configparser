Configparser is a package for reading config files.

It currently supports the following configs:

# Basic Config

Supported Data types: string, integer

Format:
```
# Comment
CONFIG1 = VALUE1
CONFIG2 = VALUE_2
```

# JSON Config

Supported Data types: string, numeric, structs, boolean, duration

Format:
```
{
	"key1": "object attribute 1",
	"key2": "object attribute 2"
}
```
