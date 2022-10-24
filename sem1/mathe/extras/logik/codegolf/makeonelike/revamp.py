import base64,sys
with open(sys.argv[1]+".oneline.py",'w') as b:
    with open(sys.argv[1],'r') as a:
        b.write('import base64;exec(base64.b64decode('+str(base64.b64encode(a.read().encode('utf-8')))+'))')