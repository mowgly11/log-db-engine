package db_operations

func Get(memtable *map[string]string, key string) string {
	return (*memtable)[key]
}