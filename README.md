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
    <td>6491 ★; 43590 commits, latest 2018-10-08</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/arangodb/arangodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/apache/couchdb">CouchDB</a></td>
    <td>CouchDB</td>
    <td>REST</td>
    <td>Erlang</td>
    <td>Apache 2.0</td>
    <td>3864 ★; 11299 commits, latest 2018-06-21</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/_/couchdb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/simonw/datasette">Datasette</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache 2.0</td>
    <td>2003 ★; 558 commits, latest 2018-10-04</td>
    <td>Read-only. <a href="https://hub.docker.com/r/terranodo/datasette/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dgraph-io/dgraph">Dgraph</a></td>
    <td>Dgraph</td>
    <td><a href="https://docs.dgraph.io/query-language/">GraphQL+-</a>, a GraphQL derivative</td>
    <td>Go</td>
    <td>Apache 2.0 with a &quot;no selling&quot; restriction</td>
    <td>6480 ★; 2607 commits, latest 2018-10-06</td>
    <td>A database with a built-in GraphQL-like API. <a href="https://hub.docker.com/r/dgraph/dgraph/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dreamfactorysoftware/dreamfactory">DreamFactory</a></td>
    <td>MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and <a href="https://www.dreamfactory.com/products">others</a>.</td>
    <td>REST</td>
    <td>PHP 5</td>
    <td>Apache 2.0, proprietary (optional extras)</td>
    <td>848 ★; 813 commits, latest 2018-09-19</td>
    <td><a href="https://hub.docker.com/r/dreamfactorysoftware/df-docker/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/pyeve/eve">Eve</a></td>
    <td>MongoDB; extensions for Elasticsearch, Neo4j, SQLAlchemy (SQL databases).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>BSD (three-clause)</td>
    <td>5253 ★; 2908 commits, latest 2018-10-04</td>
    <td>The SQLAlchemy extension isn't automatic. It requires the user to write SQLAlchemy mappings.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/hasura/graphql-engine">Hasura GraphQL Engine</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>Haskell</td>
    <td>GNU AGPLv3</td>
    <td>4108 ★; 283 commits, latest 2018-10-08</td>
    <td><a href="https://hub.docker.com/r/hasura/graphql-engine/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://bitbucket.org/prometheus/htsql/src">HTSQL</a></td>
    <td>MySQL, PostgreSQL, SQLite (free); Oracle, MS SQL (proprietary)</td>
    <td>REST</td>
    <td>Python 2</td>
    <td>GNU AGPLv3, proprietary (Oracle and MS SQL support)</td>
    <td>n/a</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/neo4j-graphql/neo4j-graphql">neo4j-graphql</a></td>
    <td>Neo4j</td>
    <td>GraphQL</td>
    <td>Kotlin</td>
    <td>Apache 2.0</td>
    <td>237 ★; 129 commits, latest 2018-05-16</td>
    <td>Can generate a GraphQL API from an existing database or derive a new database model from a GraphQL schema and auto-generate the resolvers.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/orientechnologies/orientdb">OrientDB</a></td>
    <td>OrientDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache 2.0</td>
    <td>3623 ★; 17397 commits, latest 2018-10-08</td>
    <td>A database with a built-in REST API. <a href="https://store.docker.com/images/orientdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/mevdschee/php-crud-api">PHP-CRUD-API</a></td>
    <td>MySQL, PostgreSQL, MS SQL Server.</td>
    <td>REST</td>
    <td>PHP 7</td>
    <td>MIT</td>
    <td>1729 ★; 1205 commits, latest 2018-10-08</td>
    <td>Supports GIS + automatic OpenAPI 3.0 docs.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/graphile/postgraphile">PostGraphile</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>TypeScript (Node.js)</td>
    <td>MIT</td>
    <td>5996 ★; 970 commits, latest 2018-10-08</td>
    <td>Formerly &quot;PostGraphQL&quot;, <a href="https://hub.docker.com/r/postgraphql/postgraphql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/PostgREST/postgrest">PostgREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Haskell</td>
    <td>MIT</td>
    <td>11346 ★; 1425 commits, latest 2018-09-27</td>
    <td><a href="https://hub.docker.com/r/postgrest/postgrest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prest/prest">pREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>1836 ★; 440 commits, latest 2018-09-04</td>
    <td><a href="https://hub.docker.com/r/prest/prest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/SoftInstigate/restheart">RESTHeart</a></td>
    <td>MongoDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>GNU AGPLv3</td>
    <td>494 ★; 1531 commits, latest 2018-10-08</td>
    <td><a href="https://hub.docker.com/r/softinstigate/restheart/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/jeffknupp/sandman2">sandman2</a></td>
    <td>All supported by SQLAlchemy (MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>Apache 2.0</td>
    <td>827 ★; 139 commits, latest 2018-10-05</td>
    <td></td>
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
    <td>275 ★; 52 commits, latest 2018-02-16</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/o1lab/xmysql">xmysql</a></td>
    <td>MySQL</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>3256 ★; 266 commits, latest 2018-09-28</td>
    <td><a href="https://hub.docker.com/r/markuman/xmysql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/BjoernKW/ZenQuery">ZenQuery</a></td>
    <td>PostgreSQL, MySQL, IBM Db2, Oracle Database, Microsoft SQL Server and <a href="https://github.com/BjoernKW/ZenQuery#database">others</a></td>
    <td>REST</td>
    <td>Java (JavaScript for the front-end)</td>
    <td>Apache 2.0</td>
    <td>33 ★; 282 commits, latest 2017-01-31</td>
    <td>Read-only.</td>
  </tr>
</table>


GitHub stats updated 2018-10-08. The commit count and the latest commit date are for the default branch (usually `master`).

# Related projects

For projects that depend on or enhance those on the list see the [Related projects](https://github.com/dbohdan/automatic-api/wiki/Related-projects) wiki page. Feel free to add yours.

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or to update an existing one. See [CONTRIBUTING](./CONTRIBUTING.md) for the details.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
