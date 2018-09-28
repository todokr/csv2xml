# csv2xml

Convert CSV to XML. Output is written to stdout.

## Build

```
$ go build csv2xml.go
```

# Usage

```
$ csv2xml -f source.csv
```

| option | description              | default |
|--------|--------------------------|---------|
| -f     | CSV file name to convert | -       |
| -r     | root element name of XML | root    |
| -e     | node element name of XML | item    |

If CSV is like this,

```csv
"col1","col2","col3"
"a","1","あ"
"b","2","い"
"c","3","う"
```

result is below.

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<root>
  <item>
    <col1><![CDATA[a]]></col1>
    <col2><![CDATA[1]]></col2>
    <col3><![CDATA[あ]]></col3>
  </item>
  <item>
    <col1><![CDATA[b]]></col1>
    <col2><![CDATA[2]]></col2>
    <col3><![CDATA[い]]></col3>
  </item>
  <item>
    <col1><![CDATA[c]]></col1>
    <col2><![CDATA[3]]></col2>
    <col3><![CDATA[う]]></col3>
  </item>
</root>
```

## License
MIT
