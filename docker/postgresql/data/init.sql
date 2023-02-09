SELECT 'CREATE DATABASE develop'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'develop')\gexec