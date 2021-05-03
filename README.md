# url-shortener
<h2>instructions</h2>
<ol>
  <li>
    <h3>run project</h3>
    <ol>
      <li>
        <p>setup dependencies</p>
        <code>make setup</code>
      </li>
      <li>
        <p>build application</p>
        <code>go build -o url-shortener-app app/*.go</code>
      </li>
      <li>
        <p>run application</p>
        <code>./url-shortener-app</code>
      </li>
      <li>
        <p>tear-down dependencies</p>
        <code>make tear-down</code>
      </li>
    </ol>
  </li>
  <li>
    <h3>API reference:</h3>
    <p>postman import file</p>
    <a href="https://www.getpostman.com/collections/fdb19f553547fe53d830">https://www.getpostman.com/collections/fdb19f553547fe53d830</a>
    <h3>documentation</h3>
    <pre>
client api:
└───create URL:
│   │ POST {app-url}/client/
│   │ body: {
│   │ fullURL: string eg: "http://www.google.com/"
│   │ expireDate: string eg: "2021-05-05T11:08:03Z"
│   │ }
│   │ response: created shortCode
└───get shortCode from fullURL:
    │ GET {app-url}/short-code
    │ query param: full-url
    │ eg:  {app-url}/short-code?full-url=http://www.google.com/
    │ response: shortCode
  
  
admin api: (required header "admin-token" with default value "supersecrettoken")
└───list URL:
│   │ GET {app-url}/admin/url?short-code-filter=&full-url-keyword-filter
│   │ query params: short-code-filter, full-url-keyword-filter
│   │ eg: {app-url}/admin/url?short-code-filter=&full-url-keyword-filter=google
│   │ reponse: list of URL
└───delete URL:
    │ DELETE {app-url}/admin/url/:short-code
    │ param: short-code
    │ eg: {app-url}/admin/url/0fc5sl9MR
    │ resposne: "delted"
     
redirect api:
└───redirectURL:
    GET {app-url}/redirect/:short-code
    param: short-code
    eg: {app-url}/redirect/YTGW6lrMg
    reponse: redirect to fullURL
  
</pre>

  </li>
  <li>
    <h3>edit blacklist regex</h3>
    <p>modify dev.env on variable "BLACKLIST", separate by "," eg: ^.*troll.*,^.*forbidden.*
  </li>
  <li>
    <h3>run test</h3>
    <code>go test ./...</code>
  </li>
</ol>
<h2>design</h2>
<img src="https://user-images.githubusercontent.com/21177109/116853507-60544b00-ac20-11eb-9fcd-1edda8e0a308.png"/>
