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
	"key1": "string value",
	"key2": numericValue,
	"key3": {
		"key4": "object attribute 1",
		"key5": "object attribute 2"
	}
}
```

# TOML Config 

Supported Data types: string, numeric,  boolean, duration

Format
```
[ConfigGroup1]
	Config1 = config1value
	QuotedConfig2 = "quotedValue2"
```
