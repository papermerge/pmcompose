package logs

const S3WORKER_LOGGING = `version: 1
disable_existing_loggers: false

formatters:
  verbose:
    format: '%(levelname)s %(asctime)s %(module)s %(message)s'

handlers:
  console:
    class: logging.StreamHandler
    formatter: verbose

loggers:
  celery:
    level: DEBUG
    handlers: [console]
  s3worker:
    level: DEBUG
    handlers: [console]
    propagate: no
`
