DROP DATABASE IF EXISTS "TDE_MS_UPLOAD";
CREATE DATABASE "TDE_MS_UPLOAD";
CONNECT "TDE_MS_UPLOAD";

CREATE TABLE "UPLOAD" ( 
    "user_uuid"                     UUID                UNIQUE DEFAULT gen_random_UUID(),
    "created_at"                    TIMESTAMP           DEFAULT CURRENT_TIMESTAMP
);