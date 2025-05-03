# Generate Docker Compose for Papermerge DMS


## Installation

Both `pmcompose` and `pmcompose_templates` folder must be located
in same place, as `pmcompose` looks up `pmcompose_templates` folder
relative to itself.
For example if `pmcompose` is located in `/usr/local/bin`, then `pmcompose_templates`
must be in `/usr/local/bin/pmcompose_templates`.


## Usage

To start it in interactive mode use `-i` flag:

```
$ pmcompose -i
```

You will be asked, in interactive session, about different parameters.
Docker compose file will be written in current working directory.
