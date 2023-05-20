# Templa{{.r}}
A snippet generation tool using the golang text templating language

# Usage
each object in the root array is an `operation`

each `operation` is repeated for each `value` it has

to run simply
```bash
go run .
```
<br />

`name` the template file name

`out` the output file name

`values` all the values used inside the template should be specified here

for each value 1 output file is generated

# Exmaple
`operations.json`
```json
[
    {
        "name": "example.txt",
        "out": "ex.txt",
        "values": [
            {
                "value": "test"
            },
            {
                "value": 5
            }
        ]
    }, 
] 
```

`templates/example.txt`
```txt
this is indeed an exmple value: {{.value}}
```

`output/ex/ex_1.txt`
```txt
this is indeed an exmple value: test
```

`output/ex/ex_2.txt`
```txt
this is indeed an exmple value: 5
```
