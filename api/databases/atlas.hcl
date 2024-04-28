data "external_schema" "gorm" {
  program = ["env", "ENCORERUNTIME_NOPANIC=1", "go", "run", "./scripts/atlas-gorm-loader.go"]
}

env "local" {
  src = data.external_schema.gorm.url

  migration {
    dir = "file://migrations"
    format = golang-migrate
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
