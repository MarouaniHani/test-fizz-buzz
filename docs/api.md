## TEST REST API

#### Generate strings list

* ( POST ): `/v1/generate-strings`.

1. Request

* Url:  `/v1/generate-strings`

* Body :

```
{
	"firstNumber":2,
	"secondNumber":5,
	"limit":20,
	"firstString":"hany",
	"secondString":"marouani"
}
```

1. Responses

* status: 200

```
{
	"data": [
		"1",
		"hany",
		"3",
		"hany",
		"marouani",
		"hany",
		"7",
		"hany",
		"9",
		"hanymarouani",
		"11",
		"hany",
		"13",
		"hany",
		"marouani",
		"hany",
		"17",
		"hany",
		"19",
		"hanymarouani"
	],
	"success": true
}
```