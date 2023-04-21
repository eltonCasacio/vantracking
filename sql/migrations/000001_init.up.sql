CREATE TABLE drivers(
    id varchar(36) NOT NULL PRIMARY KEY,
    cpf varchar(15) NOT NULL,
    name varchar(50) NOT NULL,
    nickname varchar(50), 
    phone varchar(15), 
    uf varchar(5), 
    city varchar(80), 
    street varchar(100), 
    number varchar(50), 
    cep varchar(10),
    complement varchar(100),
    latitude varchar(40), 
    longitude varchar(40), 
    active BOOLEAN DEFAULT true
);

CREATE TABLE monitors(
    id varchar(36) NOT NULL PRIMARY KEY,
    cpf varchar(15) NOT NULL,
    name varchar(50) NOT NULL,
    phone_number varchar(255), 
    uf varchar(5), 
    city varchar(80), 
    street varchar(100), 
    number varchar(50), 
    cep varchar(10), 
    complement varchar(100),
    latitude varchar(40), 
    longitude varchar(40), 
    active BOOLEAN DEFAULT true
);

CREATE TABLE partners(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(50) NOT NULL,
    description varchar(255) NOT NULL, 
    price varchar(20), 
    phone_number varchar(255), 
    uf varchar(5), 
    city varchar(80),
    street varchar(100),
    number varchar(50), 
    cep varchar(10), 
    complement varchar(100),
    latitude varchar(40),
    longitude varchar(40)
);

CREATE TABLE passengers(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(50) NOT NULL,
    nickname varchar(255) NOT NULL,
    route_code varchar(255), 
    goes BOOLEAN DEFAULT true, 
    comesback BOOLEAN DEFAULT true, 
    register_confirmed BOOLEAN DEFAULT false, 
    school_name varchar(100), 
    monitor_id varchar(50), 
    active BOOLEAN DEFAULT true
);

CREATE TABLE routes(
    code varchar(255) NOT NULL, 
    name varchar(50) NOT NULL, 
    driver_id varchar(255) NOT NULL, 
    started BOOLEAN DEFAULT false
);

CREATE TABLE devices(
    monitor_id varchar(255) NOT NULL PRIMARY KEY, 
    token varchar(255) NOT NULL
);
