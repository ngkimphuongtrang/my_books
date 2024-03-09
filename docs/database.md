## Context 
I want to design database for my personal book management where I track all the books I have read, there are 2 options
1. Create 2 tables: 
   - books(author, name) and 
   - reads(book_id, finished_date, source, language). 
   
   When I finish reading a new book, I create a record on books and a record on reads with created book_id on books. When I finish reading an old book, I find the book on books with known name and author, I create a record on reads with found book_id. When I need to query how many times I read for each book, I will make a join operation on 2 tables.
2. Create 1 table: 
   - reads(author, name, finished_date, source, language). 

   When I finish reading a new or an old book, I will do the same operation, insert a new row with author, name, finished_date, source, language. When I need to query how many times I read for each book, I will count on group (author, name) to have immediately result. Looks like efficient, but how I can make sure the same book will be saved in the same values in columns author and name for my query.



Both approaches have their merits, but the best choice depends on specific needs and priorities. Let's deep dive into each option's pros and cons:

### Option 1: Two Tables (Normalized Design)
Using two tables (`books` and `read_books`) is a normalized approach that prevents redundancy by separating book details from the reading records.

#### `books` Table:
- `book_id` (Primary Key)
- `author`
- `name`

#### `read_books` Table:
- `read_id` (Primary Key)
- `book_id` (Foreign Key)
- `finished_date`
- `source`
- `language`

**Pros:**
- Avoids data redundancy: Book details are stored once, regardless of how many times you read the book.
- Update efficiency: If a book's details change, you only need to update one record in the `books` table.
- Integrity: It is easier to maintain consistent data for book details.

**Cons:**
- Slightly more complex queries: To get details about the books you've read, you'll need to join tables.
- Dependency on consistent book identification: You'll need to ensure that you correctly link `read_books` entries to the `books` table.

This approach is generally preferred in database design because it is scalable and efficient in terms of storage. To ensure that we link to the correct book, we could implement a user interface that allows you to search for a book by name or author before logging it as read, reducing the chance of duplicate entries.

### Option 2: Single Table (Denormalized Design)
Keeping all data in one `reads` table simplifies the design and makes some queries more straightforward.

#### `read_books` Table:
- `read_id` (Primary Key)
- `author`
- `name`
- `finished_date`
- `source`
- `language`

**Pros:**
- Simplicity: Easier to write queries since all information is in one place.
- Immediate results: Aggregating read counts for each book is more straightforward.

**Cons:**
- Data redundancy: The author's name and book title are repeated for every reading record.
- Maintenance: Updates to a book's details require changes to multiple records.
- Potential for inconsistency: Ensuring that book information is entered consistently can be challenging.

To ensure consistency in book entries, we could use a controlled vocabulary for `author` and `name` fields or even implement a UI feature that suggests book titles as users type, based on existing entries in the database. This would help maintain consistency and prevent duplicates.

## Decision
While the single-table design might seem simpler initially, the two-table approach is more robust, scalable, and adheres to good database normalization practices. It may require more complex queries, but it will save us from potential headaches as database grows. To address the issue of ensuring consistent book identification:

- Implement a user-friendly interface that allows users to search for existing books before adding new ones.
- Use dropdowns or autocomplete features that help users select from existing authors and titles.
- Regularly review the `books` table to merge any duplicates that may have occurred.

Ultimately, if we plan to expand your database in the future or potentially integrate with other systems, the two-table approach would likely serve you better in the long run.

## References
- Chat-GPT 4.0