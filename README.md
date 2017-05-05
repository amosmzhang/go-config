# go-config

Read a yml or json config file, and automatically look for environment variables to overwrite config file settings.

See tests for examples.

## Usage

1. Create a config struct. Config struct can use types int, float64, string, bool, array, struct, as well as pointers to these types.
2. Call GetConfig() passing address of config struct and yml file name.
3. For environment variables, names should match field names. For arrays and structs, separate field names/indices with "_".

