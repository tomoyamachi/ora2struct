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

with Docker

```
$ ls .
sample.ddl

$ docker run --rm -v "$(pwd):/app" tomoyamachi/ora2struct /app/sample.ddl
$ ls .
models.go sample.ddl

$ cat models.go
package models

import (
        "database/sql"
...
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

check http://www.orafaq.com/wiki/Datapump

```sql
$ expdp scott/tiger DIRECTORY=dmpdir DUMPFILE=scott.dmp
```