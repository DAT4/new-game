import json

with open('gogame.json', 'r') as f:
    dims = json.load(f)


for x in dims['layers']:
    print('{\n\t',end="")
    y = x['data']
    for i in range(len(y)):
        if y[i] != 0:
            print(y[i]-1,end=", ")
        else:
            print(y[i],end=", ")
        if i%30 == 0 and i != 0:
            print('\n\t',end="")
    print('\n},')
