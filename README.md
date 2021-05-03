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
