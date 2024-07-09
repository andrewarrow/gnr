# gnr

curl -X POST "https://www.reddit.com/api/v1/access_token" \
    --user "key:secret" \
    -d "grant_type=password&username=%s&password=%s&scope=creddits,modcontributors,modconfig,subscribe,wikiread,wikiedit,vote,mysubreddits,submit,modlog,modposts,modflair,save,modothers,read,privatemessages,report,identity,livemanage,account,modtraffic,edit,modwiki,modself,history,flair" \
    -H "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:127.0) Gecko/20100101 Firefox/127.0"

```
#curl -H "Authorization: Token <ACCESS_TOKEN>" "https://www.oauth.reddit.com/api/comment -H "Content-Type: application/json" -d '{"parent_id": "123q", "author": "Jane Doe", "text": "This is a reply to the main parent."}'

curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "User-Agent: 'MyApp/1.0.0 (by /u/andrewfromx)'" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "thing_id=t3_12456a" \
  -d "text=I like rings." \
  https://oauth.reddit.com/api/comment
 ```
