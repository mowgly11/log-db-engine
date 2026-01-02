# Log-Structured Storage Engine - BStorage

this is a toy database engine built by me in order to apply what i learnt from the Designing Data-Intensive applications book. in chapter 3, Martin explains two families of database engines which are *log-structured enignes* and *page-oriented engines*. i've decided to start with the log-structured one since its pretty interesting how much can be put into making a simple idea such as a log to turn it into a whole database engine.

### Current Features:

1. Read a key-value pair from the database
2. Insert a key-value pair into the database
3. Hash Index for faster reads
4. Segments Manager to handle multiple data segments management
5. Compaction
