```mermaid
erDiagram
    Store {
        Id BIGSERIAL PK
        Name TEXT
        Address TEXT
        CreatedAt TIMESTAMP
        UpdatedAt TIMESTAMP
        DeletedAt TIMESTAMP
    }

    User {
        Id BIGSERIAL PK
        Name TEXT
        CreatedAt TIMESTAMP
        UpdatedAt TIMESTAMP
        DeletedAT TIMESTAMP
    }

    LaundryItem {
        Id BIGSERIAL PK
        UserId BIGINT FK
        Slug TEXT UK
        Photo TEXT
        Description TEXT
        CreatedAt TIMESTAMP
        UpdatedAt TIMESTAMP
        DeletedAT TIMESTAMP
    }

    Laundry {
        Id BIGSERIAL PK
        UserId BIGINT FK
        StoreId BIGINT FK
        Date TIMESTAMP
        PickedDate TIMESTAMP
        Price INT
        Weight FLOAT
        Type INT
        CreatedAt TIMESTAMP
        UpdatedAt TIMESTAMP
        DeletedAT TIMESTAMP
    }
    
    LaundryDetail {
        Id BIGSERIAL PK
        LaundryId BIGINT FK
        LaundryItemId BIGINT FK
        Status TEXT
        CreatedAt TIMESTAMP
        UpdatedAt TIMESTAMP
        DeletedAT TIMESTAMP
    }
    
    User || -- |{ LaundryItem : have
    Store || -- o{ Laundry : done_by
    User || -- o{ Laundry : do
    Laundry || -- o{ LaundryDetail : have
    LaundryItem || -- o{ LaundryDetail : history


```