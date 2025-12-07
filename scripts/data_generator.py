data = ""


for i in range(100, 200):
    data += f"PUT key{i}:value{i}\n"

print(data)