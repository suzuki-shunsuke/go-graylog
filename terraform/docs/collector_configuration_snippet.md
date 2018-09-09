# graylog_collector_configuration_snippet

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_collector_configuration_snippet.go

```
resource "graylog_collector_configuration_snippet" "test" {
  backend = "nxlog"
  name = "test"
  snippet = "{{if .Linux}}\nUser nxlog\nGroup nxlog\n{{if eq .LinuxPlatform \"debian\"}}\nModuledir /usr/lib/nxlog/modules\nCacheDir /var/spool/collector-sidecar/nxlog\nPidFile /var/run/graylog/collector-sidecar/nxlog.pid\n{{end}}\n{{if eq .LinuxPlatform \"redhat\"}}\nModuledir /usr/libexec/nxlog/modules\nCacheDir /var/spool/collector-sidecar/nxlog\nPidFile /var/run/graylog/collector-sidecar/nxlog.pid\n{{end}}\ndefine LOGFILE /var/log/graylog/collector-sidecar/nxlog.log\nLogFile %LOGFILE%\nLogLevel INFO\n\n<Extension logrotate>\n    Module  xm_fileop\n    <Schedule>\n        When    @daily\n        Exec    file_cycle('%LOGFILE%', 7);\n     </Schedule>\n</Extension>\n{{end}}\n{{if .Windows}}\nModuledir %ROOT%\\modules\nCacheDir %ROOT%\\data\nPidfile %ROOT%\\data\\nxlog.pid\nSpoolDir %ROOT%\\data\nLogFile %ROOT%\\data\\nxlog.log\nLogLevel INFO\n\n<Extension logrotate>\n    Module  xm_fileop\n    <Schedule>\n        When    @daily\n        Exec    file_cycle('%ROOT%\\data\\nxlog.log', 7);\n     </Schedule>\n</Extension>\n{{end}}"
  collector_configuration_id = "${graylog_collector_configuration.test-terraform.id}"
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
backend | string |
name | string |
snippet | string |
collector_configuration_id | string |

### Optional Argument

Nothing.

## Attrs Reference

name | type | etc
--- | --- | ---
snippet_id | string |
