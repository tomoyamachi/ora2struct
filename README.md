# dbscheme2struct

```sql
select
type,owner,table_name,listagg(PRIVILEGE, ',') within group (order by PRIVILEGE) as PRIVILEGEs
from dba_tab_privs
where grantee = 'xxx'
and TYPE not in ('SEQUENCE')
group by type, owner, table_name
order by type,owner,table_name;
```