# tweet-release

### Usages

```yml
name: "tweet release"
on:
  release:
    types: [published]

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - name: Tweet release
        uses: thedevsaddam/tweet-release@v1.0
        with:
          apiKey: ${{ secrets.API_KEY }}
          apiKeySecret: ${{ secrets.API_SECRET }}
          accessToken: ${{ secrets.ACCESS_TOKEN }}
          accessTokenSecret: ${{ secrets.ACCESS_TOKEN_SECRET }}
          tweet: 'Tweet release published v1.0'

```

NOTE: This is a test tweet release custom action