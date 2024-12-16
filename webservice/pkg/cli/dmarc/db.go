package dmarc

import (
	"fmt"

	"github.com/mytechpal-net/dmarc-analyzer/webservice/internal/db"
	"github.com/spf13/cobra"
)

var queryCreateTableDmarc = `CREATE TABLE IF NOT EXISTS dmarc_reports (
	id INT PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL,
	report_id VARCHAR(255) NOT NULL,
	issuer VARCHAR(255) NOT NULL,
	org_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	start_date BIGINT NOT NULL,  -- Using BIGINT for UNIX timestamp
	end_date BIGINT NOT NULL     -- Using BIGINT for UNIX timestamp
);`

var queryCreateTablePolicy = `CREATE TABLE IF NOT EXISTS dmarc_policy (
id INT PRIMARY KEY AUTOINCREMENT,
report_id INT NOT NULL,
domain VARCHAR(255) NOT NULL,
policy VARCHAR(255) NOT NULL,
subdomain_policy VARCHAR(255) NULL,
adkim VARCHAR(255) NOT NULL,
aspf VARCHAR(255) NOT NULL,
pct INT NOT NULL,
failure_option VARCHAR(255) NOT NULL,
FOREIGN KEY (report_id) REFERENCES dmarc_reports (id)
);`

var queryCreateTableSource = `CREATE TABLE IF NOT EXISTS dmarc_source (
id INT PRIMARY KEY AUTOINCREMENT,
report_id INT NOT NULL,
source_ip VARCHAR(255) NOT NULL,
count INT NOT NULL,
header_from VARCHAR(255) NOT NULL,
envelope_from VARCHAR(255) NOT NULL,
FOREIGN KEY (report_id) REFERENCES dmarc_reports (id)
);`

var queryCreateTableAuthResults = `CREATE TABLE IF NOT EXISTS dmarc_auth_results (
id INT PRIMARY KEY AUTOINCREMENT,
source_id INT NOT NULL,
spf_result VARCHAR(255) NOT NULL,
dkim_result VARCHAR(255) NOT NULL,
spf_aligned BOOLEAN NOT NULL,
dkim_aligned BOOLEAN NOT NULL,
FOREIGN KEY (source_id) REFERENCES dmarc_source (id)
);`

var queryCreateTablePolicyEvaluated = `CREATE TABLE IF NOT EXISTS dmarc_policy_evaluated (
id INT PRIMARY KEY AUTOINCREMENT,
source_id INT NOT NULL,
disposition VARCHAR(255) NOT NULL,
reason VARCHAR(255) NOT NULL,
FOREIGN KEY (source)id) REFERENCES dmarc_source (id)
);`

var DbCmd = &cobra.Command{
	Use: "db-init",
	Run: func(cmd *cobra.Command, args []string) {
		dbInit()
	},
}

func dbInit() {
	fmt.Println("Running db-init command...")
	tables := []string{queryCreateTableDmarc, queryCreateTablePolicy, queryCreateTableSource, queryCreateTableAuthResults, queryCreateTablePolicyEvaluated}

	db.Init()
	defer db.Close()

	db.CreateTables(tables)
}
