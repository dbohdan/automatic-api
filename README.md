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
    <td><a href="https://www.apinizer.com/products/api-creator/">Apinizer API Creator</a></td>
    <td>Oracle, MySQL, PostgreSQL, MsSQL, IBM DB2, SAP Sybase ASE, Apache Impala, Apache Hive</td>
    <td>REST</td>
    <td>Java</td>
    <td>Proprietary (SaaS)</td>
    <td>n/a</td>
    <td>Generates OpenAPI Specifications.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/arangodb/arangodb">ArangoDB</a></td>
    <td>ArangoDB</td>
    <td>REST</td>
    <td>C++</td>
    <td>Apache 2.0</td>
    <td>13030 ★; 50312 commits, latest 2023-07-11</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/arangodb/arangodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/apache/couchdb">CouchDB</a></td>
    <td>CouchDB</td>
    <td>REST</td>
    <td>Erlang</td>
    <td>Apache 2.0</td>
    <td>5745 ★; 13264 commits, latest 2023-07-11</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/_/couchdb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/juxt/crux">Crux</a></td>
    <td>Crux</td>
    <td>REST</td>
    <td>Clojure</td>
    <td>MIT</td>
    <td>2299 ★; 7447 commits, latest 2023-07-10</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/juxt/crux-standalone">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/simonw/datasette">Datasette</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache 2.0</td>
    <td>7969 ★; 2392 commits, latest 2023-07-08</td>
    <td>Read-only. <a href="https://hub.docker.com/r/terranodo/datasette/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dgraph-io/dgraph">Dgraph</a></td>
    <td>Dgraph</td>
    <td>GraphQL (since version 2.0.0-rc1)</td>
    <td>Go</td>
    <td>Apache 2.0</td>
    <td>19439 ★; 5985 commits, latest 2023-07-11</td>
    <td>A database with a built-in GraphQL API. <a href="https://hub.docker.com/r/dgraph/dgraph/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/directus/directus">Directus</a></td>
    <td>PostgreSQL, MySQL, SQLite, OracleDB, CockroachDB, MariaDB, MS SQL</td>
    <td>REST and GraphQL</td>
    <td>TypeScript</td>
    <td>Propretary (Business Source License 1.1), GNU GPLv3 (after three years)</td>
    <td>22356 ★; 10514 commits, latest 2023-07-11</td>
    <td><a href="https://hub.docker.com/r/directus/directus">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dreamfactorysoftware/dreamfactory">DreamFactory</a></td>
    <td>MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and <a href="https://www.dreamfactory.com/products">others</a>.</td>
    <td>REST</td>
    <td>PHP 5</td>
    <td>Apache 2.0, proprietary (optional extras)</td>
    <td>1384 ★; 1064 commits, latest 2023-03-30</td>
    <td><a href="https://hub.docker.com/r/dreamfactorysoftware/df-docker/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/Softmotions/ejdb">EJDB2</a></td>
    <td>EJDB2</td>
    <td>REST</td>
    <td>C</td>
    <td>MIT</td>
    <td>1387 ★; 2819 commits, latest 2023-03-20</td>
    <td>A database with a built-in REST API. <a href="https://github.com/Softmotions/ejdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/pyeve/eve">Eve</a></td>
    <td>MongoDB; extensions for Elasticsearch, Neo4j, SQLAlchemy (SQL databases).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>BSD (three-clause)</td>
    <td>6617 ★; 3393 commits, latest 2023-07-10</td>
    <td>The SQLAlchemy extension isn't automatic. It requires the user to write SQLAlchemy mappings.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/hasura/graphql-engine">Hasura GraphQL Engine</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>Haskell</td>
    <td>Apache 2.0</td>
    <td>29971 ★; 7430 commits, latest 2023-07-11</td>
    <td><a href="https://hub.docker.com/r/hasura/graphql-engine/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prometheusresearch/htsql">HTSQL</a></td>
    <td>MySQL, PostgreSQL, SQLite (free); Oracle, MS SQL (proprietary)</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache 2.0, proprietary (Oracle and MS SQL support)</td>
    <td>18 ★; 1235 commits, latest 2020-08-11</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/influxdata/influxdb">InfluxDB</a></td>
    <td>InfluxDB</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>25781 ★; 35272 commits, latest 2023-06-30</td>
    <td>A timeseries database with a built-in REST API. <a href="https://hub.docker.com/_/influxdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/neo4j-graphql/neo4j-graphql">neo4j-graphql</a></td>
    <td>Neo4j</td>
    <td>GraphQL</td>
    <td>Kotlin</td>
    <td>Apache 2.0</td>
    <td>450 ★; 164 commits, latest 2020-10-22</td>
    <td>Can generate a GraphQL API from an existing database or derive a new database model from a GraphQL schema and auto-generate the resolvers.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/nocodb/nocodb">NocoDB</a></td>
    <td>MySQL, PostgreSQL, SQL Server, SQLite</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>36895 ★; 15453 commits, latest 2023-07-11</td>
    <td><a href="https://hub.docker.com/r/markuman/xmysql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/orientechnologies/orientdb">OrientDB</a></td>
    <td>OrientDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache 2.0</td>
    <td>4615 ★; 21497 commits, latest 2023-07-03</td>
    <td>A database with a built-in REST API. <a href="https://store.docker.com/images/orientdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/mevdschee/php-crud-api">PHP-CRUD-API</a></td>
    <td>MySQL, PostgreSQL, MS SQL Server.</td>
    <td>REST</td>
    <td>PHP 7</td>
    <td>MIT</td>
    <td>3361 ★; 2097 commits, latest 2023-07-05</td>
    <td>Supports GIS + automatic OpenAPI 3.0 docs.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/graphile/postgraphile">PostGraphile</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>TypeScript (Node.js)</td>
    <td>MIT</td>
    <td>12035 ★; 1398 commits, latest 2023-06-22</td>
    <td>Formerly &quot;PostGraphQL&quot;, <a href="https://hub.docker.com/r/postgraphql/postgraphql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/PostgREST/postgrest">PostgREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Haskell</td>
    <td>MIT</td>
    <td>20776 ★; 2483 commits, latest 2023-07-08</td>
    <td><a href="https://hub.docker.com/r/postgrest/postgrest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prest/prest">pREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>3846 ★; 1933 commits, latest 2023-06-03</td>
    <td><a href="https://hub.docker.com/r/prest/prest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/geekypedia/pRESTige">pRESTige</a></td>
    <td>MySQL</td>
    <td>REST</td>
    <td>PHP</td>
    <td>MIT</td>
    <td>96 ★; 1081 commits, latest 2022-07-13</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/SoftInstigate/restheart">RESTHeart</a></td>
    <td>MongoDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>GNU AGPLv3</td>
    <td>766 ★; 4144 commits, latest 2023-06-20</td>
    <td><a href="https://hub.docker.com/r/softinstigate/restheart/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/restsql/restsql">restSQL</a></td>
    <td>MySQL, PostgreSQL</td>
    <td>REST</td>
    <td>Java</td>
    <td>MIT</td>
    <td>136 ★; 54 commits, latest 2018-10-18</td>
    <td><a href="https://hub.docker.com/r/restsql/service/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/jeffknupp/sandman2">sandman2</a></td>
    <td>All supported by SQLAlchemy (MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>Apache 2.0</td>
    <td>1953 ★; 250 commits, latest 2020-12-21</td>
    <td><a href="https://hub.docker.com/r/jeffknupp/sandman2/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/siodb/siodb">Siodb</a></td>
    <td>Siodb</td>
    <td>REST</td>
    <td>C++, Go</td>
    <td>GNU AGPLv3</td>
    <td>42 ★; 278 commits, latest 2023-02-18</td>
    <td>A database with a built-in REST API. <a href="https://hub.docker.com/r/siodb/siodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/thevahidal/soul">Soul</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>813 ★; 106 commits, latest 2023-07-02</td>
    <td>A SQLite RESTful server</td>
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
    <td><a href="https://github.com/supabase/supabase">Supabase</a></td>
    <td>PostgreSQL</td>
    <td>REST and GraphQL</td>
    <td>TypeScript, Elixir (Realtime), Rust (pg_graphql), Go (GoTrue)</td>
    <td>Apache 2.0</td>
    <td>53035 ★; 22027 commits, latest 2023-07-11</td>
    <td>Uses PostgREST. <a href="https://supabase.com/docs/guides/self-hosting/docker">Guide to self-Hosting with Docker</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/bradleyboy/tuql">tuql</a></td>
    <td>SQLite 3 or SQL infile</td>
    <td>GraphQL</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>1013 ★; 72 commits, latest 2021-06-06</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/nicolasff/webdis">Webdis</a></td>
    <td>Redis</td>
    <td>REST</td>
    <td>C</td>
    <td>BSD (two-clause)</td>
    <td>2697 ★; 629 commits, latest 2023-06-14</td>
    <td>Supports pub/sub with chunked transfer encoding and WebSockets.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/wundergraph/wundergraph">WunderGraph</a></td>
    <td>PostgreSQL, MySQL, SQLite, SQL Server, MongoDB + Atlas, PlanetScale, YugabyteDB, Neon.tech, FaunaDB</td>
    <td>GraphQL</td>
    <td>TypeScript, Go</td>
    <td>Apache 2.0</td>
    <td>1755 ★; 1664 commits, latest 2023-07-11</td>
    <td><a href="https://docs.wundergraph.com/">WunderGraph Docs</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/BjoernKW/ZenQuery">ZenQuery</a></td>
    <td>PostgreSQL, MySQL, IBM Db2, Oracle Database, Microsoft SQL Server and <a href="https://github.com/BjoernKW/ZenQuery#database">others</a></td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache 2.0</td>
    <td>63 ★; 283 commits, latest 2018-10-16</td>
    <td>Read-only.</td>
  </tr>
</table>


GitHub stats updated 2023-07-11. The commit count and the latest commit date are for the default branch (usually `master`).

# Related projects

For projects that depend on or enhance those on the list see the [Related projects](https://github.com/dbohdan/automatic-api/wiki/Related-projects) wiki page. Feel free to add yours.

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or to update an existing one. See [CONTRIBUTING](./CONTRIBUTING.md) for the details.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
