-- *************** SqlDBM: PostgreSQL ****************;
-- ***************************************************;


-- ************************************** "Users"

CREATE TABLE IF NOT EXISTS "Users"
(
 "userid"           serial NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_Users" PRIMARY KEY ( "userid" )
);








-- ************************************** "Networks"

CREATE TABLE IF NOT EXISTS "Networks"
(
 "networkid"        serial NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_Networks" PRIMARY KEY ( "networkid" )
);








-- ************************************** "Hosts"

CREATE TABLE IF NOT EXISTS "Hosts"
(
 "hostid"           serial NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_Hosts" PRIMARY KEY ( "hostid" )
);








-- ************************************** "Groups"

CREATE TABLE IF NOT EXISTS "Groups"
(
 "groupid"          serial NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_Groups" PRIMARY KEY ( "groupid" )
);








-- ************************************** "Domains"

CREATE TABLE IF NOT EXISTS "Domains"
(
 "domainid"         serial NOT NULL,
 "currentTimestamp" timestamp NULL,
 CONSTRAINT "PK_Domains" PRIMARY KEY ( "domainid" )
);








-- ************************************** "UserDetails"

CREATE TABLE IF NOT EXISTS "UserDetails"
(
 "timestamp"    timestamp NOT NULL,
 "userid"       integer NOT NULL,
 "endtimestamp" timestamp NULL,
 "username"     text NOT NULL,
 "password"     text NOT NULL,
 CONSTRAINT "PK_table_166" PRIMARY KEY ( "timestamp", "userid" ),
 CONSTRAINT "Ind_244" UNIQUE ( "username" ),
 CONSTRAINT "FK_183" FOREIGN KEY ( "userid" ) REFERENCES "Users" ( "userid" )
);

CREATE INDEX "fkIdx_183" ON "UserDetails"
(
 "userid"
);








-- ************************************** "NetworkPermissions"

CREATE TABLE IF NOT EXISTS "NetworkPermissions"
(
 "groupid"          integer NOT NULL,
 "networkid"        integer NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_NetworkPermission" PRIMARY KEY ( "groupid", "networkid" ),
 CONSTRAINT "FK_112" FOREIGN KEY ( "groupid" ) REFERENCES "Groups" ( "groupid" ),
 CONSTRAINT "FK_87" FOREIGN KEY ( "networkid" ) REFERENCES "Networks" ( "networkid" )
);

CREATE INDEX "fkIdx_112" ON "NetworkPermissions"
(
 "groupid"
);

CREATE INDEX "fkIdx_87" ON "NetworkPermissions"
(
 "networkid"
);








-- ************************************** "NetworkMemberships"

CREATE TABLE IF NOT EXISTS "NetworkMemberships"
(
 "hostid"           integer NOT NULL,
 "networkid"        integer NOT NULL,
 "currentTimestamp" timestamp NULL,
 CONSTRAINT "PK_NetworkMembership" PRIMARY KEY ( "hostid", "networkid" ),
 CONSTRAINT "FK_44" FOREIGN KEY ( "hostid" ) REFERENCES "Hosts" ( "hostid" ),
 CONSTRAINT "FK_51" FOREIGN KEY ( "networkid" ) REFERENCES "Networks" ( "networkid" )
);

CREATE INDEX "fkIdx_44" ON "NetworkMemberships"
(
 "hostid"
);

CREATE INDEX "fkIdx_51" ON "NetworkMemberships"
(
 "networkid"
);








-- ************************************** "NetworkDetails"

CREATE TABLE IF NOT EXISTS "NetworkDetails"
(
 "timestamp"      timestamp NOT NULL,
 "networkid"      integer NOT NULL,
 "endtimestamp"   timestamp NULL,
 "networkADDR"    inet NOT NULL,
 "dhcprangeStart" inet NOT NULL,
 "dhcprangeStop"  inet NOT NULL,
 "description"    text NOT NULL,
 "cidr"           bit (64) NOT NULL,
 CONSTRAINT "PK_NetworkHistory" PRIMARY KEY ( "timestamp", "networkid" ),
 CONSTRAINT "FK_238" FOREIGN KEY ( "networkid" ) REFERENCES "Networks" ( "networkid" )
);

CREATE INDEX "fkIdx_238" ON "NetworkDetails"
(
 "networkid"
);








-- ************************************** "HostDetails"

CREATE TABLE IF NOT EXISTS "HostDetails"
(
 "timestamp"   timestamp NOT NULL,
 "hostid"      integer NOT NULL,
 "mac"         bit(48) NOT NULL,
 "description" text NULL,
 CONSTRAINT "PK_HostHistory" PRIMARY KEY ( "timestamp", "hostid" ),
 CONSTRAINT "Ind_242" UNIQUE ( "mac" ),
 CONSTRAINT "FK_126" FOREIGN KEY ( "hostid" ) REFERENCES "Hosts" ( "hostid" )
);

CREATE INDEX "fkIdx_126" ON "HostDetails"
(
 "hostid"
);








-- ************************************** "GroupMemberships"

CREATE TABLE IF NOT EXISTS "GroupMemberships"
(
 "groupid"          integer NOT NULL,
 "userid"           integer NOT NULL,
 "currenttimestamp" timestamp NULL,
 CONSTRAINT "PK_GroupMembership" PRIMARY KEY ( "groupid", "userid" ),
 CONSTRAINT "FK_103" FOREIGN KEY ( "groupid" ) REFERENCES "Groups" ( "groupid" ),
 CONSTRAINT "FK_119" FOREIGN KEY ( "userid" ) REFERENCES "Users" ( "userid" )
);

CREATE INDEX "fkIdx_103" ON "GroupMemberships"
(
 "groupid"
);

CREATE INDEX "fkIdx_119" ON "GroupMemberships"
(
 "userid"
);








-- ************************************** "GroupDetails"

CREATE TABLE IF NOT EXISTS "GroupDetails"
(
 "timestamp"   timestamp NOT NULL,
 "groupid"     integer NOT NULL,
 "name"        varchar(50) NOT NULL,
 "description" text NULL,
 CONSTRAINT "PK_GroupHistory" PRIMARY KEY ( "timestamp", "groupid" ),
 CONSTRAINT "Ind_243" UNIQUE ( "name" ),
 CONSTRAINT "FK_198" FOREIGN KEY ( "groupid" ) REFERENCES "Groups" ( "groupid" )
);

CREATE INDEX "fkIdx_198" ON "GroupDetails"
(
 "groupid"
);








-- ************************************** "DomainPermissions"

CREATE TABLE IF NOT EXISTS "DomainPermissions"
(
 "domainid"         integer NOT NULL,
 "groupid"          integer NOT NULL,
 "currentTimestamp" timestamp NULL,
 CONSTRAINT "PK_DomainPermission" PRIMARY KEY ( "domainid", "groupid" ),
 CONSTRAINT "FK_116" FOREIGN KEY ( "groupid" ) REFERENCES "Groups" ( "groupid" ),
 CONSTRAINT "FK_79" FOREIGN KEY ( "domainid" ) REFERENCES "Domains" ( "domainid" )
);

CREATE INDEX "fkIdx_116" ON "DomainPermissions"
(
 "groupid"
);

CREATE INDEX "fkIdx_79" ON "DomainPermissions"
(
 "domainid"
);








-- ************************************** "DomainMemberships"

CREATE TABLE IF NOT EXISTS "DomainMemberships"
(
 "domainid"         integer NOT NULL,
 "hostid"           integer NOT NULL,
 "currentTimestamp" timestamp NULL,
 CONSTRAINT "PK_DomainMembership" PRIMARY KEY ( "domainid", "hostid" ),
 CONSTRAINT "FK_37" FOREIGN KEY ( "hostid" ) REFERENCES "Hosts" ( "hostid" ),
 CONSTRAINT "FK_40" FOREIGN KEY ( "domainid" ) REFERENCES "Domains" ( "domainid" )
);

CREATE INDEX "fkIdx_37" ON "DomainMemberships"
(
 "hostid"
);

CREATE INDEX "fkIdx_40" ON "DomainMemberships"
(
 "domainid"
);








-- ************************************** "DomainDetails"

CREATE TABLE IF NOT EXISTS "DomainDetails"
(
 "timestamp"    timestamp NOT NULL,
 "domainid"     integer NOT NULL,
 "endtimestamp" timestamp NULL,
 "name"         varchar(50) NOT NULL,
 "description"  text NULL,
 CONSTRAINT "PK_DomainHistory" PRIMARY KEY ( "timestamp", "domainid" ),
 CONSTRAINT "FK_187" FOREIGN KEY ( "domainid" ) REFERENCES "Domains" ( "domainid" )
);

CREATE INDEX "fkIdx_187" ON "DomainDetails"
(
 "domainid"
);








-- ************************************** "NetworkPermissionsDetails"

CREATE TABLE IF NOT EXISTS "NetworkPermissionsDetails"
(
 "timestamp"    timestamp NOT NULL,
 "groupid"      integer NOT NULL,
 "networkid"    integer NOT NULL,
 "endtimestamp" timestamp NULL,
 CONSTRAINT "PK_NetworkPermissionsHistory" PRIMARY KEY ( "timestamp", "groupid", "networkid" ),
 CONSTRAINT "FK_168" FOREIGN KEY ( "groupid", "networkid" ) REFERENCES "NetworkPermissions" ( "groupid", "networkid" )
);

CREATE INDEX "fkIdx_168" ON "NetworkPermissionsDetails"
(
 "groupid",
 "networkid"
);








-- ************************************** "NetworkMembershipDetails"

CREATE TABLE IF NOT EXISTS "NetworkMembershipDetails"
(
 "timestamp"    timestamp NOT NULL,
 "hostid"       integer NOT NULL,
 "networkid"    integer NOT NULL,
 "endtimestamp" timestamp NULL,
 "address"      inet NULL,
 "dynamic"      boolean NOT NULL,
 CONSTRAINT "PK_HostHistory" PRIMARY KEY ( "timestamp", "hostid", "networkid" ),
 CONSTRAINT "FK_140" FOREIGN KEY ( "hostid", "networkid" ) REFERENCES "NetworkMemberships" ( "hostid", "networkid" )
);

CREATE INDEX "fkIdx_140" ON "NetworkMembershipDetails"
(
 "hostid",
 "networkid"
);








-- ************************************** "GroupPermissionDetails"

CREATE TABLE IF NOT EXISTS "GroupPermissionDetails"
(
 "timestamp"    timestamp NOT NULL,
 "groupid"      integer NOT NULL,
 "userid"       integer NOT NULL,
 "endtimestamp" timestamp NULL,
 CONSTRAINT "PK_GroupPermissionHistory" PRIMARY KEY ( "timestamp", "groupid", "userid" ),
 CONSTRAINT "FK_178" FOREIGN KEY ( "groupid", "userid" ) REFERENCES "GroupMemberships" ( "groupid", "userid" )
);

CREATE INDEX "fkIdx_178" ON "GroupPermissionDetails"
(
 "groupid",
 "userid"
);








-- ************************************** "DomainPermissionDetails"

CREATE TABLE IF NOT EXISTS "DomainPermissionDetails"
(
 "timestamp"    timestamp NOT NULL,
 "domainid"     integer NOT NULL,
 "groupid"      integer NOT NULL,
 "endtimestamp" timestamp NULL,
 CONSTRAINT "PK_DomainPermissionHistory" PRIMARY KEY ( "timestamp", "domainid", "groupid" ),
 CONSTRAINT "FK_173" FOREIGN KEY ( "domainid", "groupid" ) REFERENCES "DomainPermissions" ( "domainid", "groupid" )
);

CREATE INDEX "fkIdx_173" ON "DomainPermissionDetails"
(
 "domainid",
 "groupid"
);








-- ************************************** "DomainMembershipDetails"

CREATE TABLE IF NOT EXISTS "DomainMembershipDetails"
(
 "timestamp"      timestamp NOT NULL,
 "domainid"       integer NOT NULL,
 "hostid"         integer NOT NULL,
 "endtimestamp"   timestamp NULL,
 "hostname"       varchar(50) NOT NULL,
 "generateRecord" boolean NOT NULL,
 CONSTRAINT "PK_DomainMembershipHistory" PRIMARY KEY ( "timestamp", "domainid", "hostid" ),
 CONSTRAINT "FK_148" FOREIGN KEY ( "domainid", "hostid" ) REFERENCES "DomainMemberships" ( "domainid", "hostid" )
);

CREATE INDEX "fkIdx_148" ON "DomainMembershipDetails"
(
 "domainid",
 "hostid"
);







