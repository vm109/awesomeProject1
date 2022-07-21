### altering table

```sql
    ALTER TABLE <table_name> ADD <column_name> <data_type>
``` 

### alter table add foregin key 

```sql
    ALTER TABLE <table_name> ADD FOREIGN KEY <column_name> REFERENCES <other_table_name>(<other_table_column>)
```

```sql
    ALTER TABLE  <table_name> ADD PRIMARY KEY (<column_name>)
```

```sql
    ALTER TABLE <table_name> ADD CONSTRAINT <primary_key_name> PRIMARY KEY (<column_name1, column_name2>)
```

### drop trigger
```sql
 DROP TRIGGER [IF EXISTS] <trigger_name> on <table_name> [CASCADE|RESTRICT]
```