file = open("database.txt", "a")
data = ""

for i in range(10):
    for j in range(1000 * i, 1000 * i+1):
        data += f"key{j}:value{j}"
        file.write(data)
        data = ""
