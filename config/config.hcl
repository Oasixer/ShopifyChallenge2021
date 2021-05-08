// Mode: The mode that the server runs in
// can be prod, dev, or test
Mode = "dev"

// Port: The port it runs on
// this only matters in dev
Port = "3000"

// DB: the block to hold all database values
DB  {
    Host="localhost"
    Port="5432"
    User="postgres"
    Password="postgres"
    DBName="shopify"
    SSLMode="disable"
}
