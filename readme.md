
clean architecture
- PORT is the interface that connect to the outer layer


definition
- adapter คือ data ที่เชื่อมต่อ
- port คือ interface ที่เอา adapter (data) มาใช้งาน
reposity
this is about the reposity that connect to DB directly

service
this is about the service that connect to user


first create repository folder
1. create customer.go first to define data schema that associated with customer database (create type Customer struct)
2. create type CustomerRepository interface that define method to get data from DB

create service folder
1. create customer.go to define interface of usage and adapter of customer data // doing as the same as repository but this one is for service