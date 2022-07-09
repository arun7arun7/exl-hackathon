CREATE DATABASE IF NOT EXISTS exl;

-- CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'root';
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;

CREATE USER IF NOT EXISTS 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON *.* TO 'docker'@'%' WITH GRANT OPTION;

USE exl;

CREATE TABLE IF NOT EXISTS azure_tenant 
(tenant_id VARCHAR(255) PRIMARY KEY, storage_account VARCHAR(255), 
container_name VARCHAR(255), client_id VARCHAR(255), client_secret VARCHAR(255));

CREATE TABLE IF NOT EXISTS files
(object_id VARCHAR(255) PRIMARY KEY, file_extension VARCHAR(255), 
tenant_id VARCHAR(255), cloud_type VARCHAR(255));

INSERT IGNORE INTO azure_tenant(tenant_id, storage_account, container_name, client_id, client_secret)
VALUES("de18e1ee-5536-4959-961c-bcfb59c93e26", "prac", "exl", "1ff1fc50-373f-4a9d-9553-631b58f7e97a", "_7u8Q~Pzz8-tt2DNeX_zvNYGZN~.efEoVJ~tgbfK");
