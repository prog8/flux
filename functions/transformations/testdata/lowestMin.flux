from(bucket:"test")
    |> range(start: 2018-11-07T00:00:00Z)
    |> lowestMin(n: 3, by: ["_measurement", "host"])
