### Issue and Challenges:

1. Universal character recognition, right now the db only supports ASCII characters with limitations (the "**:**" breaks the engine)
2. Optimize the compaction algorithm
3. Fix a case when inserting json objects ("{"title": "hello"}" in the case the program reads the "**:"** as a separator)
