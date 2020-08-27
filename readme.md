# En API til at få et overblik
Man kan få et overblik over studerende og deres spørgsmål med denne api.

## Studerende

For at få listen over alle studerende:

```
http://127.0.0.1:8080/?query={list{firstName,lastName,id,discord,middleName}}
```

For at få en enkelt studerendes info ved at søge ID

```
http://127.0.0.1:8080/?query={student(id="s195469"){firstName,lastName,id,discord,middleName}}
```

## Spørgsmål

For at se alle spørgsmål:

```
http://127.0.0.1:8080/?query={questions{student{id,firstName,middleName,lastName,discord,role},question,channelId,topic,active}}
```

Dette vil give et svar lignende:

```json
 {
  "data": {
    "questions": [
      {
        "active": true,
        "channelId": "748525251889070100",
        "question": "Hvilket lokale skal vi mødes i den første dag??",
        "student": {
          "discord": "623123053940834354",
          "firstName": "Martin",
          "id": "s195469",
          "lastName": "Mårtensson",
          "middleName": "",
          "role": "TA"
        },
        "topic": [
          "studiestart"
        ]
      }
    ]
  }
}
```
