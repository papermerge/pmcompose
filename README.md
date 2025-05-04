# Generate Docker Compose for Papermerge DMS


## Installation

Both `pmcompose` and `pmcompose_templates` folder must be located
in same place, as `pmcompose` looks up `pmcompose_templates` folder
relative to itself.
For example if `pmcompose` is located in `/usr/local/bin`, then `pmcompose_templates`
must be in `/usr/local/bin/pmcompose_templates`.


1. Download latest release
2. unzip pmcompose_linux_amd64.zip
3. then run following commands:

```
sudo mv pmcompose /usr/local/bin/
sudo mv pmcompose_templates /usr/local/bin/
```


## Usage

To start it in interactive mode use `-i` flag:

```
$ pmcompose -i
```

You will be asked, in interactive session, about different parameters.
Docker compose file will be written in current working directory.


To generate basic docker compose for Papermerge 3.5 non-iteractively use:

```
$ pmcompose -u admin -p pass123
```

To generate docker compose with logging configs:

```
$ pmcompose -u admin -p pass123 -lc
```

To generate docker compose with s3 backend:

```
$ pmcompose -u admin -p pass123 -lc -s3
```

For help use `-h` flag:

```
$ pmcompose -h
```
