-- CREATE roles Table with id,name,is_active
CREATE TABLE IF NOT EXISTS roles
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE
);

-- CREATE permissions Table IF NOT EXISTS with id,name,is_active
CREATE TABLE IF NOT EXISTS permissions
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE
);

-- CREATE role_permissions Table IF NOT EXISTS with id,role_id,permission_id
CREATE TABLE IF NOT EXISTS role_permissions
(
    role_id     INT NOT NULL,
    permission_id INT NOT NULL
);

-- CREATE user_roles Table IF NOT EXISTS with id,user_id,role_id
CREATE TABLE IF NOT EXISTS user_roles
(
    user_id     INT NOT NULL,
    role_id     INT NOT NULL
);