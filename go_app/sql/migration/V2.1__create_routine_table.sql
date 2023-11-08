CREATE TABLE IF NOT EXISTS routines
 (
    id varchar(32) PRIMARY KEY NOT NULL,
    routine_name varchar(255), 
    created_at timestamp, 
    completion_status float, 
    -- foreign key references to dailytasks table for collection of daily tasks.
    daily_task_id varchar(26) references dailytasks(id)
 );

 -- Now lets create a routine_name unique validation trigger

 CREATE OR REPLACE FUNCTION check_if_routine_name_unique()
 RETURNS TRIGGER AS $$
 BEGIN 
  IF EXISTS (
    SELECT 1 FROM routines WHERE routine_name = NEW.routine_name LIMIT 1
  ) THEN 
   RAISE EXCEPTION 'routine_name must be unique';
  END IF;
  RETURN NEW;
 END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER routines_unique_routine_name
BEFORE INSERT OR UPDATE ON routines
FOR EACH ROW EXECUTE FUNCTION check_if_routine_name_unique();