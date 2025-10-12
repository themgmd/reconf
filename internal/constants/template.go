package constants

import "strings"

const DefaultConfigTemplate = `##
##  File: values.yaml
##  description: предназначен для описания конфигурации приложения
##  значения задаются в двух блоках
##  config: значения не содержащие пароли и другую секрутную информацию
##  secret: по ключу задаются ключи волта, в котором хранятся переменные
##
##  в файлах values_development, values_staging, values_products - можно
##  переопределить значение переменных конфигурации для конкретного окружения
##

config:
  # http
  http_port: 8080

  # grpc
  grpc_port: 8090

  # postgres
  postgres_max_conns: 50
  postgres_min_conns: 10

secret:
  # postgres
  postgres_dsn: pg_dsn`

const environmentConfigTemplate = `##
##  File: values_{{.Env}}.yaml
##  description: предназначен для описания конфигурации приложения в {{.Env}} окружении
##  значения задаются в двух блоках
##  config: значения не содержащие пароли и другую секрутную информацию
##  secret: по ключу задаются ключи волта, в котором хранятся переменные
##

config:
  # http
  http_port: 6080

  # grpc
  grpc_port: 6090

secret:
  # postgres
  postgres_dsn: {{.Env}}_postgres_dsn
`

func EnvironmentConfigTemplate(envName string) string {
	return strings.ReplaceAll(environmentConfigTemplate, "{{.Env}}", envName)
}
