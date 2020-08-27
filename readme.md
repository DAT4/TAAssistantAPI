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
