CREATE DATABASE IF NOT EXISTS exl;

CREATE USER IF NOT EXISTS 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON *.* TO 'docker'@'%' WITH GRANT OPTION;

USE exl;

CREATE TABLE IF NOT EXISTS azure_tenant 
(organization_id VARCHAR(255) PRIMARY KEY, tenant_id VARCHAR(255), storage_account VARCHAR(255), 
container_name VARCHAR(255), client_id VARCHAR(255), client_secret VARCHAR(255));

CREATE TABLE IF NOT EXISTS aws_tenant
(organization_id VARCHAR(255) PRIMARY KEY, aws_region VARCHAR(255), bucket_name VARCHAR(255), 
access_key_id VARCHAR(255), secret_access_key VARCHAR(255));

CREATE TABLE IF NOT EXISTS files
(object_id VARCHAR(255) PRIMARY KEY, file_extension VARCHAR(255), 
organization_id VARCHAR(255), cloud_type VARCHAR(255));

INSERT IGNORE INTO azure_tenant(organization_id, tenant_id, storage_account, container_name, client_id, client_secret)
VALUES("exl-client-1", "de18e1ee-5536-4959-961c-bcfb59c93e26", "prac", "exl", "1ff1fc50-373f-4a9d-9553-631b58f7e97a", "_7u8Q~Pzz8-tt2DNeX_zvNYGZN~.efEoVJ~tgbfK");

INSERT IGNORE INTO aws_tenant(organization_id, aws_region, bucket_name, access_key_id, secret_access_key)
VALUES("exl-client-1", "us-east-2", "exl-storage", "AKIA52EZ53Y4GJP6CEX2", "J2gcn5aqq/rsJul6uJeWT06lNL10ZtwubJWoMG82");
