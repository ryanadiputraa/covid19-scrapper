# covid19-scrapper

simple REST API that serve covid case data from covid19.go.id

---

## API SPEC

### Global Cases

-   method: `GET`
-   endpoint: `/api/covidcases`
-   header:
    -   Content-Type: `application/json`
    -   Accept: `application/json`
-   body:

```json
{
    "code": 200,
    "data": {
        "totalCountry": 162,
        "confirmedCased": 119611553,
        "deaths": 2605995
    }
}
```

---

### Indonesia Cases

-   method: `GET`
-   endpoint: `/api/covidcases`
-   header:
    -   Content-Type: `application/json`
    -   Accept: `application/json`
-   body:

```json
{
    "code": 200,
    "data": {
        "positive": 1894025,
        "recover": 1735144,
        "deaths": 52566
    }
}
```
