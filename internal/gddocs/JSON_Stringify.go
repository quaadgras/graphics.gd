/*
## JSON.stringify(my_dictionary)
{"name":"my_dictionary","version":"1.0.0","entities":[{"name":"entity_0","value":"value_0"},{"name":"entity_1","value":"value_1"}]}

## JSON.stringify(my_dictionary, "\t")
{
	"name": "my_dictionary",
	"version": "1.0.0",
	"entities": [
		{
			"name": "entity_0",
			"value": "value_0"
		},
		{
			"name": "entity_1",
			"value": "value_1"
		}
	]
}

## JSON.stringify(my_dictionary, "...")
{
..."name": "my_dictionary",
..."version": "1.0.0",
..."entities": [
......{
........."name": "entity_0",
........."value": "value_0"
......},
......{
........."name": "entity_1",
........."value": "value_1"
......}
...]
}
*/

package main
