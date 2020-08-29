import requests
import json

data = {
        'query':'{student(id:"s195469"){id}}',
        }

#url = 'http://localhost:8080/?query=mutation+_{answer(id:123,answer:"Dette er svaret",student:"s195469"){answer{answer}}}'

query = '''
http://localhost:8080/?query=
query+_{
    question(id:1598562170){
        id
        timestamp
        question
        answer{
            answer
        }
    }
}
'''
mutation = '''
http://localhost:8080/?query=
mutation+_{
    answer(id:1598562170, answer:"Vi skal mødes i bygning 303, i lokale 44", student:"s195469"){
        question
        answer{
            answer
        }
    }
}
'''
#r = requests.get(query)
r = requests.get(mutation)
print(json.dumps(r.json(),indent=2,ensure_ascii=False))
