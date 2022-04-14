window.BENCHMARK_DATA = {
  "lastUpdate": 1649898150477,
  "repoUrl": "https://github.com/posteris/database",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "11edf845da4829e458b8e6a452708db437824cb7",
          "message": "Organize the benchmark",
          "timestamp": "2022-04-13T21:52:02-03:00",
          "tree_id": "e632dad55541a18cef874a970fddae0f6e0dfacf",
          "url": "https://github.com/posteris/database/commit/11edf845da4829e458b8e6a452708db437824cb7"
        },
        "date": 1649898149713,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConnection/PostgreSQL",
            "value": 0.005617,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkConnection/SQLite",
            "value": 0.000271,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkConnection/MySQL",
            "value": 0.0009766,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}