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
         "ph":10,
         "tds":10,
         "ox":10,
         "tmp":25,
         "wqi":30,
         "location": {
            "latitude": 0.1234,
            "longitude": 0.43214
        }
      },
      {
         "date":"1996-10-15T00:05:32.000Z",
         "ph":10,
         "tds":10,
         "ox":10,
         "tmp":25,
         "wqi":30,
         "location": {
            "latitude": 0.1234,
            "longitude": 0.43214
        }
      }
   ]
}
```


#### /reports/rivers/{river_id}
##### Estado general de un río: el promedio actual de todos sus módulos
```
{
    "wqi": {
        "min": 10,
        "max": 10,
        "average": 10,
        "stdDeviation": 10,
    },
    "ph": {
        "min": 10,
        "max": 10,
        "average": 10,
        "stdDeviation": 10,
    },
    "tds": {
        "min": 10,
        "max": 10,
        "average": 10,
        "stdDeviation": 10,
    },
    "do": {
        "min": 10,
        "max": 10,
        "average": 10,
        "stdDeviation": 10,
    },
    "tmp": {
        "min": 10,
        "max": 10,
        "average": 10,
        "stdDeviation": 10,
    },
}
```

/reports/rivers/{river_id}/details
<!-- /reports/rivers/{river_id}/modules -->

