# Log Based Storage Engine - BStorage

### Current Features:

1. Read a key-value pair from the database
2. Insert a key-value pair into the database
3. Hash Index for faster reads
4. Segments Manager to handle multiple data segments management
5. Compaction

### Issue and Challenges:

1. Universal character recognition, right now the db only supports ASCII characters with limitations (the : is vulenrable)
2. file naming should be exact, and the script doesn't check for useless files, this needs to be fixed (possible solution is to save the timestamp in the file name)
