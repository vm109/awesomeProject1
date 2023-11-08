#!/bin/sh

MIGRATION_DIR=$(cd .. && cd .. && pwd)/sql/migration

# Check if schema is already baselined
SCHEMA_ALREADY_BASELINED=$(docker run --rm -v ${MIGRATION_DIR}:/flyway/sql 026088843893.dkr.ecr.us-east-1.amazonaws.com/flyway/flyway:9.2 \
    -url=jdbc:postgresql://db.reqpokkztpjowzlugnng.supabase.co:5432/postgres \
    -user=postgres -password="BH09gf2430#!" info | grep "no migration yet" | wc -l)

# Baselining the schema only if not already baselined
if [ "$SCHEMA_ALREADY_BASELINED" -eq 1 ]; then
    echo "Schema is already baselined. Skipping baseline."
else
    docker run --rm -v ${MIGRATION_DIR}:/flyway/sql 026088843893.dkr.ecr.us-east-1.amazonaws.com/flyway/flyway:9.2 \
        -url=jdbc:postgresql://db.reqpokkztpjowzlugnng.supabase.co:5432/postgres \
        -user=postgres -password="BH09gf2430#!" baseline
fi

# Applying migrations
docker run --rm -v ${MIGRATION_DIR}:/flyway/sql 026088843893.dkr.ecr.us-east-1.amazonaws.com/flyway/flyway:9.2 \
    -url=jdbc:postgresql://db.reqpokkztpjowzlugnng.supabase.co:5432/postgres \
    -user=postgres -password="BH09gf2430#!" -outOfOrder=true migrate
