window.BENCHMARK_DATA = {
  "lastUpdate": 1649885163446,
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
          "id": "f1b9bd2edd8753ff53ab65bd04ff718e87c37e28",
          "message": "add postgres and mysql to Performance test",
          "timestamp": "2022-04-13T18:09:51-03:00",
          "tree_id": "c114bb6295506b7090e06682431db34241f95531",
          "url": "https://github.com/posteris/database/commit/f1b9bd2edd8753ff53ab65bd04ff718e87c37e28"
        },
        "date": 1649884301640,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGetAllowedDB/all-databases",
            "value": 0.0000044,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/none",
            "value": 0.0000041,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/postgres",
            "value": 0.0000038,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/sqlite",
            "value": 0.0000058,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/mysql",
            "value": 0.0000038,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/unknown",
            "value": 0.000005,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/postgres-successfull",
            "value": 0.005738,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/postgres-error",
            "value": 0.00528,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/sqlite-successfull",
            "value": 0.0002902,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/sqlite-error",
            "value": 0.0001236,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/mysql-successfull",
            "value": 0.001096,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/mysql-error",
            "value": 0.0000452,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/postgres-successfull",
            "value": 0.05232,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/sqlite-successfull",
            "value": 0.002436,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/mysql-successfull",
            "value": 0.01062,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      },
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
          "id": "51db69be0951d1b3244e63e2817b04710f0d9945",
          "message": "create full CI/CD pipeline",
          "timestamp": "2022-04-13T18:19:39-03:00",
          "tree_id": "b99904de42256e015fac2c4a5afd1a1d291eb49f",
          "url": "https://github.com/posteris/database/commit/51db69be0951d1b3244e63e2817b04710f0d9945"
        },
        "date": 1649885163002,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGetAllowedDB/all-databases",
            "value": 0.000004,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/none",
            "value": 0.0000038,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/postgres",
            "value": 0.000004,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/sqlite",
            "value": 0.0000047,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_getDatabaseType/mysql",
            "value": 0.0000043,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/unknown",
            "value": 0.0000055,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/postgres-successfull",
            "value": 0.006486,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/postgres-error",
            "value": 0.006732,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/sqlite-successfull",
            "value": 0.0003335,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/sqlite-error",
            "value": 0.0000677,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/mysql-successfull",
            "value": 0.001314,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew/mysql-error",
            "value": 0.0000479,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/postgres-successfull",
            "value": 0.04921,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/sqlite-successfull",
            "value": 0.002311,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkNew_with_migrations/mysql-successfull",
            "value": 0.01091,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}