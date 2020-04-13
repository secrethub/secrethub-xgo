from secrethub import client
import sys

c = client.Client()
result = c.read(sys.argv[1])
print(result)