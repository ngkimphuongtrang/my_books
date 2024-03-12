# my_books

## Context
### Motivation

I track the books I read throughout the year using a Note app, but tallying the annual total becomes a manual and time-consuming task. 

To streamline this process, I have created a personal book management application that not only stores a record of the books I've read but also generates statistics effortlessly, such as my yearly reading totals.

### Requirements
#### Functional

- User story 1. As a user, I need the ability to create a book entry through a simple request.
- User story 2. As a user, I need to log my reading sessions, detailing when and how I read each book.
- User story 3. As an avid reader with a vast library, I require a search function to verify whether I have previously read a specific book.
- User story 4. As a user, I want to easily view statistics, such as the number of books I've read in a given year.

#### Non-functional

- Consistency: The server should provide clear and instructive responses, ensuring understandability even when a user submits an erroneous request.

## Decision
Given the context provided, it would be prudent to introduce a new module called "My Books," which would be responsible for creating book entries, logging reading sessions, and querying insights about books and reading trends in the database.

### Database design
Read [here](docs/database.md) to see how this design was decided.
```sql
Table books {
  id integer [primary key ]
  name varchar
  author varchar 
  created_at timestamp 
  updated_at timestamp
}

Table reads {
  id integer [primary key]
  book_id integer 
  source varchar
  language varchar
  finished_date timestamp 
  created_at timestamp
  updated_at timestamp 
}

Ref: reads.book_id > books.id 
```
![Schema](docs/db-diagram.png)

### APIs 

#### Create a book
<details> 
    <summary><code>POST</code><code><b>/books</b></code></summary>
Tracker creates a book with this information

##### Body
| Name   | Required | Type   | Description            |
|--------|----------|--------|------------------------|
| name   | Y        | string | Name of the book       |
| author | Y        | string | Author of the the book |

##### Response 
| Status Code | Verdict           | Body                | Description                            |
|-------------|-------------------|---------------------|----------------------------------------|
| 200         | success           | `"data": {"id": 7}` | Success, Return id of the created book |
| 400         | invalid_parameter |                     |                                        |

##### Example 
- cURL

- Response

```json
{
  "data": {
    "id": 7
  },
  "message": "book is created successfully",
  "time": "2024-03-09T15:04:12+07:00",
  "verdict": "success"
}
```
</details>

#### Get list of books
<details> 
    <summary><code>GET</code><code><b>/books</b></code></summary>

##### Parameters
| Name     | Required | Type   | Description                                           |
|----------|----------|--------|-------------------------------------------------------|
| page     |          | int    | The page number of the results to fetch, default: 1   |
| per_page |          | int    | The number of results per page (max 100), default: 30 |
| search   |          | string | The key string to search on book name                 |

##### Response
| Status Code | Verdict           | Body                                                                                                                                                                       | Description                    |
|-------------|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------|
| 200         | success           | `"data": {"count": 7,"items": [{ "id": 1,"name": "Giết con chim nhại","author": "","created_at": "2024-03-08T20:05:58+07:00","updated_at": "2024-03-08T20:05:58+07:00"}]}` | Success, Return  list of books |
| 400         | invalid_parameter |                                                                                                                                                                            |                                |

##### Example
- cURL

- Response

```json 
{
    "data": {
        "count": 7,
        "items": [
            {
                "id": 1,
                "name": "Giết con chim nhại",
                "author": "",
                "created_at": "2024-03-08T20:05:58+07:00",
                "updated_at": "2024-03-08T20:05:58+07:00"
            }
        ]
    },
    "message": "get list of books successfully",
    "time": "2024-03-09T15:30:05+07:00",
    "verdict": "success"
}
```
</details>

#### Create a read 
<details> 
    <summary><code>POST</code><code><b>/reads</b></code></summary>

Tracker creates a read with a created book

##### Body
| Name          | Required | Type      | Description                                          |
|---------------|----------|-----------|------------------------------------------------------|
| book_id       | Y        | int       | ID of the created book you have just finished read   |
| source        | Y        | string    | Source of book you read: hard_copy, soft_copy, audio |
| language      | Y        | string    | Language of the book you read, example: EN, VI       |
| finished_date | Y        | timestamp | Date you finish reading the book                     |

##### Response
| Status Code | Verdict   | Body                 | Description                        |
|-------------|-----------|----------------------|------------------------------------|
| 200         | success   | `"data": {"id": 7}`  | Success, Return ID of created read |
| 404         | not_found |                      |                                    |

</details>

#### Get list of reads
<details> 
    <summary><code>GET</code><code><b>/reads</b></code></summary>

##### Parameters
| Name      | Required | Type    | Description                                                             |
|-----------|----------|---------|-------------------------------------------------------------------------|
| page      |          | int     | The page number of the results to fetch, default: 1                     |
| per_page  |          | int     | The number of results per page (max 100), default: 30                   |
| from_year |          | string  | The start year to search on, example: 2012                              |
| to_year   |          | string  | The end year to search on, example: 2012                                |
| language  |          | string  | The language to search on, example: VI                                  |
| source    |          | string  | The source to search on, limited on values: hard_copy, soft_copy, audio |

##### Response
| Status Code | Verdict           | Body | Description                    |
|-------------|-------------------|------|--------------------------------|
| 200         | success           |      | Success, Return  list of reads |
| 400         | invalid_parameter |      |                                |
</details>

