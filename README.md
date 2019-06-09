# ChelZone
Backend server and API for ChelZone, written in golang. The backend uses the web framework Gin to set up API routes to be accessed by a frontend, written in Vue.js.

## Resources
- https://gitlab.com/dword4/nhlapi/blob/master/stats-api.md
  - Use game by game content via http://statsapi.web.nhl.com/api/v1/game/2018021002/content; 2018021002=gameID
- https://github.com/gin-gonic/gin
- https://developer.okta.com/blog/2018/10/23/build-a-single-page-app-with-go-and-vue
- https://developers.google.com/youtube/v3/quickstart/go
  - The following link will provide activity for a particular channel:
    - https://www.googleapis.com/youtube/v3/activities?part=contentDetails,snippet&channelId=UCkUjSzthJUlO0uyUpiJfnxg&key={YOUR_API_KEY}

## MVP

### Team Stats
Each team will have its own page for stats such as team stats (including standings), and game video content. For game content, search the json object for the following top level topics:
- media
  - milestones: contains moments in the game that are described their title (type of event) ie. Period Start, Goal, Hit, Shot
- copyright
- highlights

#### To-Do List
- In the milestones section, retrieve all events that are titled "Goal" then download their highest quality video:
    - JSON object: media.milestones.items[Index of Item with title "Goal"].highlight
      - title
      - blurb
      - description
      - duration
    - Video: media.milestones.items[Index of Item with title "Goal"].highlight.playbacks[Index of item with title "FLASH_1800K_960X540"].url
