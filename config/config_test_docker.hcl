// Mode: The mode that the server runs in
// can be prod, dev, or test
Mode = "test"

// Port: The port it runs on
Port = "3000"

// DB: the block to hold all database values
DB {
    Host="db"
    Port="5434"
    User="postgres"
    Password="postgres"
    DBName="postgres"
    SSLMode="disable"
}
