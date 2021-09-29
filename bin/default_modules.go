package bin

import (
	"github.com/effluxio/zgrab2"
	"github.com/effluxio/zgrab2/modules"
	"github.com/effluxio/zgrab2/modules/bacnet"
	"github.com/effluxio/zgrab2/modules/banner"
	"github.com/effluxio/zgrab2/modules/dnp3"
	"github.com/effluxio/zgrab2/modules/fox"
	"github.com/effluxio/zgrab2/modules/ftp"
	"github.com/effluxio/zgrab2/modules/http"
	"github.com/effluxio/zgrab2/modules/imap"
	"github.com/effluxio/zgrab2/modules/ipp"
	"github.com/effluxio/zgrab2/modules/modbus"
	"github.com/effluxio/zgrab2/modules/mongodb"
	"github.com/effluxio/zgrab2/modules/mssql"
	"github.com/effluxio/zgrab2/modules/mysql"
	"github.com/effluxio/zgrab2/modules/ntp"
	"github.com/effluxio/zgrab2/modules/oracle"
	"github.com/effluxio/zgrab2/modules/pop3"
	"github.com/effluxio/zgrab2/modules/postgres"
	"github.com/effluxio/zgrab2/modules/redis"
	"github.com/effluxio/zgrab2/modules/siemens"
	"github.com/effluxio/zgrab2/modules/smb"
	"github.com/effluxio/zgrab2/modules/smtp"
	"github.com/effluxio/zgrab2/modules/telnet"
)

var defaultModules zgrab2.ModuleSet

func init() {
	defaultModules = map[string]zgrab2.ScanModule{
		"bacnet":   &bacnet.Module{},
		"banner":   &banner.Module{},
		"dnp3":     &dnp3.Module{},
		"fox":      &fox.Module{},
		"ftp":      &ftp.Module{},
		"http":     &http.Module{},
		"imap":     &imap.Module{},
		"ipp":      &ipp.Module{},
		"modbus":   &modbus.Module{},
		"mongodb":  &mongodb.Module{},
		"mssql":    &mssql.Module{},
		"mysql":    &mysql.Module{},
		"ntp":      &ntp.Module{},
		"oracle":   &oracle.Module{},
		"pop3":     &pop3.Module{},
		"postgres": &postgres.Module{},
		"redis":    &redis.Module{},
		"siemens":  &siemens.Module{},
		"smb":      &smb.Module{},
		"smtp":     &smtp.Module{},
		"ssh":      &modules.SSHModule{},
		"telnet":   &telnet.Module{},
		"tls":      &modules.TLSModule{},
	}
}

// NewModuleSetWithDefaults returns a newly allocated ModuleSet containing all
// ScanModules implemented by the ZGrab2 framework.
func NewModuleSetWithDefaults() zgrab2.ModuleSet {
	out := zgrab2.ModuleSet{}
	defaultModules.CopyInto(out)
	return out
}
