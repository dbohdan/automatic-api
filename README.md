The following is a list of tools that automatically expose a REST, GraphQL, or another kind of API for your database, as well as databases with a built-in HTTP API.

<table>
  <tr>
    <th>Project name/link</th>
    <th>Database(s) supported</th>
    <th>API type</th>
    <th>Implementation language</th>
    <th>License</th>
    <th>GitHub stats</th>
    <th>Notes</th>
  </tr>
  <tr>
    <td><a href="https://github.com/arangodb/arangodb">ArangoDB</a></td>
    <td>ArangoDB</td>
    <td>REST</td>
    <td>C++</td>
    <td>Apache 2.0</td>
    <td>9602 ★; 46388 commits, latest 2020-05-29</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/arangodb/arangodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/apache/couchdb">CouchDB</a></td>
    <td>CouchDB</td>
    <td>REST</td>
    <td>Erlang</td>
    <td>Apache 2.0</td>
    <td>4556 ★; 12323 commits, latest 2020-05-30</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/_/couchdb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/simonw/datasette">Datasette</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache 2.0</td>
    <td>3462 ★; 1040 commits, latest 2020-06-01</td>
    <td>Read-only. <a href="https://hub.docker.com/r/terranodo/datasette/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dgraph-io/dgraph">Dgraph</a></td>
    <td>Dgraph</td>
    <td>GraphQL (since version 2.0.0-rc1)</td>
    <td>Go</td>
    <td>Apache 2.0</td>
    <td>13279 ★; 4486 commits, latest 2020-06-01</td>
    <td>A database with a built-in GraphQL API. <a href="https://hub.docker.com/r/dgraph/dgraph/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dreamfactorysoftware/dreamfactory">DreamFactory</a></td>
    <td>MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and <a href="https://www.dreamfactory.com/products">others</a>.</td>
    <td>REST</td>
    <td>PHP 5</td>
    <td>Apache 2.0, proprietary (optional extras)</td>
    <td>1078 ★; 948 commits, latest 2020-05-18</td>
    <td><a href="https://hub.docker.com/r/dreamfactorysoftware/df-docker/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/pyeve/eve">Eve</a></td>
    <td>MongoDB; extensions for Elasticsearch, Neo4j, SQLAlchemy (SQL databases).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>BSD (three-clause)</td>
    <td>6105 ★; 3228 commits, latest 2020-05-16</td>
    <td>The SQLAlchemy extension isn't automatic. It requires the user to write SQLAlchemy mappings.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/hasura/graphql-engine">Hasura GraphQL Engine</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>Haskell</td>
    <td>GNU AGPLv3</td>
    <td>16967 ★; 1845 commits, latest 2020-05-29</td>
    <td><a href="https://hub.docker.com/r/hasura/graphql-engine/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prometheusresearch/htsql">HTSQL</a></td>
    <td>MySQL, PostgreSQL, SQLite (free); Oracle, MS SQL (proprietary)</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache 2.0, proprietary (Oracle and MS SQL support)</td>
    <td>3 ★; 1234 commits, latest 2020-03-16</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/neo4j-graphql/neo4j-graphql">neo4j-graphql</a></td>
    <td>Neo4j</td>
    <td>GraphQL</td>
    <td>Kotlin</td>
    <td>Apache 2.0</td>
    <td>402 ★; 163 commits, latest 2020-03-18</td>
    <td>Can generate a GraphQL API from an existing database or derive a new database model from a GraphQL schema and auto-generate the resolvers.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/orientechnologies/orientdb">OrientDB</a></td>
    <td>OrientDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache 2.0</td>
    <td>4150 ★; 19352 commits, latest 2020-05-30</td>
    <td>A database with a built-in REST API. <a href="https://store.docker.com/images/orientdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/mevdschee/php-crud-api">PHP-CRUD-API</a></td>
    <td>MySQL, PostgreSQL, MS SQL Server.</td>
    <td>REST</td>
    <td>PHP 7</td>
    <td>MIT</td>
    <td>2436 ★; 1732 commits, latest 2020-05-26</td>
    <td>Supports GIS + automatic OpenAPI 3.0 docs.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/graphile/postgraphile">PostGraphile</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>TypeScript (Node.js)</td>
    <td>MIT</td>
    <td>8967 ★; 1210 commits, latest 2020-05-12</td>
    <td>Formerly &quot;PostGraphQL&quot;, <a href="https://hub.docker.com/r/postgraphql/postgraphql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/PostgREST/postgrest">PostgREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Haskell</td>
    <td>MIT</td>
    <td>14458 ★; 1653 commits, latest 2020-05-28</td>
    <td><a href="https://hub.docker.com/r/postgrest/postgrest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prest/prest">pREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>2288 ★; 443 commits, latest 2020-04-16</td>
    <td><a href="https://hub.docker.com/r/prest/prest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/geekypedia/pRESTige">pRESTige</a></td>
    <td>MySQL</td>
    <td>REST</td>
    <td>PHP</td>
    <td>MIT</td>
    <td>49 ★; 1027 commits, latest 2020-04-29</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/SoftInstigate/restheart">RESTHeart</a></td>
    <td>MongoDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>GNU AGPLv3</td>
    <td>602 ★; 3025 commits, latest 2020-05-28</td>
    <td><a href="https://hub.docker.com/r/softinstigate/restheart/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/jeffknupp/sandman2">sandman2</a></td>
    <td>All supported by SQLAlchemy (MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>Apache 2.0</td>
    <td>1501 ★; 224 commits, latest 2020-05-19</td>
    <td><a href="https://hub.docker.com/r/jeffknupp/sandman2/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://subzero.cloud">subZero</a></td>
    <td>PostgreSQL</td>
    <td>REST and GraphQL</td>
    <td>Haskell, Lua</td>
    <td>Proprietary</td>
    <td>n/a</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/bradleyboy/tuql">tuql</a></td>
    <td>SQLite 3 or SQL infile</td>
    <td>GraphQL</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>460 ★; 67 commits, latest 2019-09-14</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/nicolasff/webdis">Webdis</a></td>
    <td>Redis</td>
    <td>REST</td>
    <td>C</td>
    <td>BSD (two-clause)</td>
    <td>2254 ★; 439 commits, latest 2020-06-01</td>
    <td>Supports pub/sub with chunked transfer encoding and WebSockets.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/o1lab/xmysql">xmysql</a></td>
    <td>MySQL</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>4085 ★; 288 commits, latest 2020-03-26</td>
    <td><a href="https://hub.docker.com/r/markuman/xmysql/">Official Docker image</a>.</td>
  </tr>
    <tr>
    <td><a href="https://github.com/xgenecloud/xgenecloud">XgeneCloud</a></td>
    <td>PostgreSQL, MySQL, SQL Server, SQLite, MariaDb</td>
    <td>REST and GraphQL</td>
    <td>JavaScript (Node.js)</td>
    <td>Apache 2.0</td>
    <td>770 ★; 26 commits, latest 2020-06-12</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/BjoernKW/ZenQuery">ZenQuery</a></td>
    <td>PostgreSQL, MySQL, IBM Db2, Oracle Database, Microsoft SQL Server and <a href="https://github.com/BjoernKW/ZenQuery#database">others</a></td>
    <td>REST</td>
    <td>Java (JavaScript for the front-end)</td>
    <td>Apache 2.0</td>
    <td>47 ★; 283 commits, latest 2018-10-16</td>
    <td>Read-only.</td>
  </tr>
</table>


GitHub stats updated 2020-06-01. The commit count and the latest commit date are for the default branch (usually `master`).

# Related projects

For projects that depend on or enhance those on the list see the [Related projects](https://github.com/dbohdan/automatic-api/wiki/Related-projects) wiki page. Feel free to add yours.

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or to update an existing one. See [CONTRIBUTING](./CONTRIBUTING.md) for the details.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
