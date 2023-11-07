### Go Routines With Channels

- There are few 100 documents in a database for a organization in a multi-tenant environment.
- *Now lets say the documents are not indexed based on date and also the database does not have sort function on date.*
- Now that org wants to delete all those documents which are prior to a date assuming documents are saved into database with `created_at` with dd-MM-YYY format.
