SELECT 'CREATE DATABASE develop'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'develop')\gexec
SELECT 'CREATE DATABASE test'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'test')\gexec