package logs

const WEBAPP_LOGGING = `version: 1
disable_existing_loggers: false

formatters:
  verbose:
    format: '%(levelname)s %(asctime)s %(module)s %(message)s'

handlers:
  console:
    class: logging.StreamHandler
    formatter: verbose

loggers:
  papermerge.core.tasks:
    level: DEBUG
    handlers: [console]
    propagate: no
`
