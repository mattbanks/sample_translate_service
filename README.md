# Sample Translate Service

Go microservice to give a REST endpoint for passing in text strings and getting language translations.

Accepts a JSON object and returns it back for now, with an appended timestamp.

### Request

Passed to `https://mighty-garden-67967.herokuapp.com/translate`

```json
{
    "id": 123,
    "language": "en",
    "content": "some content here"
}
```

### Response:

```json
{
  "id": 123,
  "language": "en",
  "content": "some content here",
  "translated_on": "2016-07-29T14:44:46.674855169-07:00"
}
```
