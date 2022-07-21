## Triggers
What is a trigger?
- A trigger is a stored procedure in database which automatically invokes whenever a defined event occurs in the database. 

What is Syntax for trigger? 
```sql
create trigger <name> [before|after] {insert|update|delete} on <table_name>
[for each row]
[trigger_body]
```

example: 
```sql
create function total_greetings ()
    returns trigger language plpgsql as $$
    begin
    select count(*)                                                                                    
    end;                                                                                                
end $$;
```

Can you create a trigger without function? 
In postgressql we cannot create trigger without creating a procedure function.

Trigger functions cannot have declared arguments
```sql
    CREATE FUNCTION <function_name>()
        returns trigger
        language plpgsql
        as
        $$
        DECLARE
            variables;
        BEGIN
            <SQL STATEMENT>
        END;
        $$;            
```

Create a Function which returns Trigger: 

```sql
 CREATE OR REPLACE FUNCTION idCounter()                         
     RETURNS TRIGGER LANGUAGE plpgsql AS $$ 
     DECLARE counter integer; plus_counter integer;
         BEGIN 
             counter:= count(*) from movie_cast;
             plus_counter:= counter+1;
             NEW._id:= CONCAT('cast', plus_counter); 
             RETURN NEW;
         END;
    $$;
```