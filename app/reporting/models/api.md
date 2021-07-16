#### /reports/modules?riverID=RVR1234?cardinality=hour
```
[
    {
        "moduleID": "MDL1234",
        "riverID": "RVR1234",
        "wqi": 10,
        "ph": 10,
        "tds": 10,
        "do": 10,
        "tmp": 10,
        "lastUpdated": "1996-10-15T00:05:32.000Z"
        "location": {
            "latitude": 0.1234,
            "longitude": 0.43214
        }
    }
]
```

#### /reports/modules/{module_id}
##### Estado general de un módulo, el último registro que se tiene de ese módulo. O debería ser el promedio hasta entonces?

Respuesta:
```
{
    "wqi": 10,
    "ph": 10,
    "tds": 10,
    "do": 10,
    "tmp": 10,
    "lastUpdated": "1996-10-15T00:05:32.000Z"
}
```

#### /reports/modules/{module_id}/details?cardinality
##### Detalles de un módulo (todos los measurements en el rango de fecha indicada)
```
{
   "moduleID":"MDL1234",
   "riverID": "RVR1234",
   "count":2,
   "minWQI": 10,
   "maxWQI": 20,
   "averageWQI":10,
   "stdDeviation":10,
   "daysCovered": [0,1]
   "data":[
      {
         "date":"1996-10-15T00:05:32.000Z",
         "parameters": [
             {
                "name": "ph",
                "value": 10
             },
             {
                "name": "tds",
                "value": 10
             },
             {
                "name": "tmp",
                "value": 10
             },
             {
                "name": "do",
                "value": 10
             },
             {
                "name": "wqi",
                "value": 10
             },
         ],
         "location": {
            "latitude": 0.1234,
            "longitude": 0.43214
        }
      },
      {
         "date":"1996-10-15T00:05:32.000Z",
         "parameters": [
             {
                "name": "ph",
                "value": 10
             },
             {
                "name": "tds",
                "value": 10
             },
             {
                "name": "tmp",
                "value": 10
             },
             {
                "name": "do",
                "value": 10
             },
             {
                "name": "wqi",
                "value": 10
             },
         ],
         "location": {
            "latitude": 0.1234,
            "longitude": 0.43214
        }
      }
   ]
}
```


#### /reports/rivers/{river_id}
##### Estado general de un río: el promedio actual(última medición) de todos sus módulos
```
{    
   "riverID": RVR123",
    "wqi": 10,
      "data":[
      {
         "date":"1996-10-15T00:05:32.000Z",
         "parameters": [
             {
                "name": "ph",
                "value": 10
             },
             {
                "name": "tds",
                "value": 10
             },
             {
                "name": "tmp",
                "value": 10
             },
             {
                "name": "do",
                "value": 10
             },
             {
                "name": "wqi",
                "value": 10
             },
         ]
      }
   ]
}
```

/reports/rivers/{river_id}/details
<!-- /reports/rivers/{river_id}/modules -->

