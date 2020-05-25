# ora2struct

a DDL file of Oracle Database convert to go struct suited [jmoiron/sqlx](https://github.com/jmoiron/sqlx).

## Quick Start

```
$ ora2struct -o /path/to/output.go /path/to/sample.ddl
$ cat output.go
package models

import (
        "database/sql"
        "time"
)

// Table SCOTT.EMP
type Emp struct {
        Empno sql.NullInt64 `db:"EMPNO"` // type: NUMBER
        Ename sql.NullString `db:"ENAME"` // type: VARCHAR2
        Job sql.NullString `db:"JOB"` // type: VARCHAR2
        Mgr sql.NullInt64 `db:"MGR"` // type: NUMBER
        Hiredate sql.NullTime `db:"HIREDATE"` // type: DATE
        Hiredaten time.Time `db:"HIREDATEN"` // type: DATE
        Sal sql.NullInt64 `db:"SAL"` // type: NUMBER
        Comm sql.NullInt64 `db:"COMM"` // type: NUMBER
        Deptno sql.NullInt64 `db:"DEPTNO"` // type: NUMBER
}
```

### Options

```
  --template value, -t value  use template
  --output value, -o value    output file name (default: "models.go")
  --package value, -p value   export package name (default: "models")
  --debug, -d                 debug mode (default: false)
  --help, -h                  show help (default: false)
```

## Exporting tables and views from a Oracle Database

```sql
select
type,owner,table_name,listagg(PRIVILEGE, ',') within group (order by PRIVILEGE) as PRIVILEGEs
from dba_tab_privs
where grantee = 'xxx'
and TYPE not in ('SEQUENCE')
group by type, owner, table_name
order by type,owner,table_name;
```