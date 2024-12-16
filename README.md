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
    <td>Apache-2.0</td>
    <td>13616 ★; 51743 commits, latest 2024-12-20</td>
    <td>A database with a built-in REST API.
<a href="https://hub.docker.com/r/arangodb/arangodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/apache/couchdb">CouchDB</a></td>
    <td>CouchDB</td>
    <td>REST</td>
    <td>Erlang</td>
    <td>Apache-2.0</td>
    <td>6312 ★; 13835 commits, latest 2024-12-19</td>
    <td>A database with a built-in REST API.
<a href="https://hub.docker.com/r/_/couchdb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/simonw/datasette">Datasette</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache-2.0</td>
    <td>9645 ★; 2663 commits, latest 2024-11-29</td>
    <td>Read-only.
<a href="https://hub.docker.com/r/terranodo/datasette/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/kdhrubo/db2rest">DB2Rest</a></td>
    <td>PostgreSQL, MySQL, MariaDB, Oracle, CockroachDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache-2.0</td>
    <td>246 ★; 1363 commits, latest 2024-12-15</td>
    <td><a href="https://hub.docker.com/r/kdhrubo/db2rest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dgraph-io/dgraph">Dgraph</a></td>
    <td>Dgraph</td>
    <td>GraphQL (since version 2.0.0-rc1)</td>
    <td>Go</td>
    <td>Apache-2.0</td>
    <td>20537 ★; 6228 commits, latest 2024-12-19</td>
    <td>A database with a built-in GraphQL API.
<a href="https://hub.docker.com/r/dgraph/dgraph/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/directus/directus">Directus</a></td>
    <td>PostgreSQL, MySQL, SQLite, OracleDB, CockroachDB, MariaDB, MS SQL</td>
    <td>REST and GraphQL</td>
    <td>TypeScript</td>
    <td>Propretary (BUSL-1.1), GPL-3.0 (after three years)</td>
    <td>28460 ★; 12461 commits, latest 2024-12-20</td>
    <td><a href="https://hub.docker.com/r/directus/directus">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dreamfactorysoftware/dreamfactory">DreamFactory</a></td>
    <td>MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and
<a href="https://www.dreamfactory.com/products">others</a>.</td>
    <td>REST</td>
    <td>PHP 5</td>
    <td>Apache-2.0, proprietary (optional extras)</td>
    <td>1574 ★; 1140 commits, latest 2024-05-16</td>
    <td><a href="https://hub.docker.com/r/dreamfactorysoftware/df-docker/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/Softmotions/ejdb">EJDB2</a></td>
    <td>EJDB2</td>
    <td>REST</td>
    <td>C</td>
    <td>MIT</td>
    <td>1445 ★; 2849 commits, latest 2024-12-02</td>
    <td>A database with a built-in REST API.
<a href="https://github.com/Softmotions/ejdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/pyeve/eve">Eve</a></td>
    <td>MongoDB; extensions for Elasticsearch, Neo4j, SQLAlchemy (SQL databases).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>BSD-3-Clause</td>
    <td>6716 ★; 3403 commits, latest 2024-10-15</td>
    <td>The SQLAlchemy extension isn't automatic.
It requires the user to write SQLAlchemy mappings.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/dosco/graphjin">GraphJin Service</a></td>
    <td>PostgreSQL, MySQL, Yugabyte</td>
    <td>GraphQL</td>
    <td>Go</td>
    <td>Apache-2.0</td>
    <td>2939 ★; 936 commits, latest 2024-09-06</td>
    <td><a href="https://graphjin.com/posts/service">Using GraphJin as a standlone service</a>.
<a href="https://hub.docker.com/r/dosco/graphjin/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/ardatan/graphql-mesh">GraphQL Mesh</a></td>
    <td>PostgreSQL (PostGraphile), MongoDB (graphql-compose-mongoose), SQLite 3 (tuql), Neo4j (Neo4j GraphQL Library), MySQL</td>
    <td>GraphQL</td>
    <td>TypeScript (Node.js)</td>
    <td>MIT</td>
    <td>3310 ★; 7995 commits, latest 2024-12-18</td>
    <td>Provides a common GraphQL gateway for different APIs and databases.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/hasura/graphql-engine">Hasura GraphQL Engine</a></td>
    <td>PostgreSQL, MS SQL, MySQL, Snowflake, and <a href="https://hasura.io/connectors">others</a></td>
    <td>GraphQL, REST</td>
    <td>Haskell</td>
    <td>Apache-2.0</td>
    <td>31246 ★; 9118 commits, latest 2024-12-20</td>
    <td><a href="https://hub.docker.com/r/hasura/graphql-engine/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prometheusresearch/htsql">HTSQL</a></td>
    <td>MySQL, PostgreSQL, SQLite (free); Oracle, MS SQL (proprietary)</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>Apache-2.0, proprietary (Oracle and MS SQL support)</td>
    <td>25 ★; 1235 commits, latest 2020-08-11</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/influxdata/influxdb">InfluxDB</a></td>
    <td>InfluxDB</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>29182 ★; 49422 commits, latest 2024-12-20</td>
    <td>A timeseries database with a built-in REST API.
<a href="https://hub.docker.com/_/influxdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/neo4j-graphql/neo4j-graphql">neo4j-graphql</a></td>
    <td>Neo4j</td>
    <td>GraphQL</td>
    <td>Kotlin</td>
    <td>Apache-2.0</td>
    <td>449 ★; 164 commits, latest 2020-10-22</td>
    <td>Can generate a GraphQL API from an existing database
or derive a new database model from a GraphQL schema and auto-generate the resolvers.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/nocodb/nocodb">NocoDB</a></td>
    <td>MySQL, PostgreSQL, SQL Server, SQLite</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>50363 ★; 25435 commits, latest 2024-12-20</td>
    <td><a href="https://hub.docker.com/r/markuman/xmysql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/orientechnologies/orientdb">OrientDB</a></td>
    <td>OrientDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache-2.0</td>
    <td>4765 ★; 25921 commits, latest 2024-12-19</td>
    <td>A database with a built-in REST API.
<a href="https://store.docker.com/images/orientdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/mevdschee/php-crud-api">PHP-CRUD-API</a></td>
    <td>MySQL, PostgreSQL, MS SQL Server.</td>
    <td>REST</td>
    <td>PHP 7</td>
    <td>MIT</td>
    <td>3628 ★; 2154 commits, latest 2024-11-22</td>
    <td>Supports GIS + automatic OpenAPI 3.0 docs.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/graphile/postgraphile">PostGraphile</a></td>
    <td>PostgreSQL</td>
    <td>GraphQL</td>
    <td>TypeScript (Node.js)</td>
    <td>MIT</td>
    <td>12635 ★; 11435 commits, latest 2024-12-19</td>
    <td>Formerly &quot;PostGraphQL&quot;.
<a href="https://hub.docker.com/r/postgraphql/postgraphql/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/PostgREST/postgrest">PostgREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Haskell</td>
    <td>MIT</td>
    <td>24100 ★; 3843 commits, latest 2024-12-18</td>
    <td><a href="https://hub.docker.com/r/postgrest/postgrest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/prest/prest">pREST</a></td>
    <td>PostgreSQL</td>
    <td>REST</td>
    <td>Go</td>
    <td>MIT</td>
    <td>4254 ★; 1998 commits, latest 2024-11-28</td>
    <td><a href="https://hub.docker.com/r/prest/prest/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/geekypedia/pRESTige">pRESTige</a></td>
    <td>MySQL</td>
    <td>REST</td>
    <td>PHP</td>
    <td>MIT</td>
    <td>108 ★; 1089 commits, latest 2024-05-31</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/SoftInstigate/restheart">RESTHeart</a></td>
    <td>MongoDB</td>
    <td>REST</td>
    <td>Java</td>
    <td>AGPL-3.0</td>
    <td>813 ★; 4958 commits, latest 2024-12-13</td>
    <td><a href="https://hub.docker.com/r/softinstigate/restheart/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/restsql/restsql">restSQL</a></td>
    <td>MySQL, PostgreSQL</td>
    <td>REST</td>
    <td>Java</td>
    <td>MIT</td>
    <td>147 ★; 54 commits, latest 2018-10-18</td>
    <td><a href="https://hub.docker.com/r/restsql/service/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/jeffknupp/sandman2">sandman2</a></td>
    <td>All supported by SQLAlchemy
(MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).</td>
    <td>REST</td>
    <td>Python 2/3</td>
    <td>Apache-2.0</td>
    <td>2011 ★; 250 commits, latest 2020-12-21</td>
    <td><a href="https://hub.docker.com/r/jeffknupp/sandman2/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/DEAL-US/Silence">Silence</a></td>
    <td>MySQL, MariaDB</td>
    <td>REST</td>
    <td>Python 3</td>
    <td>MIT</td>
    <td>17 ★; 445 commits, latest 2024-02-14</td>
    <td>Automatically generates REST endpoints to do CRUD operations against the database,
test studs, and JS files to access the API.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/siodb/siodb">Siodb</a></td>
    <td>Siodb</td>
    <td>REST</td>
    <td>C++, Go</td>
    <td>AGPL-3.0</td>
    <td>45 ★; 278 commits, latest 2023-02-18</td>
    <td>A database with a built-in REST API.
<a href="https://hub.docker.com/r/siodb/siodb/">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/thevahidal/soul">Soul</a></td>
    <td>SQLite 3</td>
    <td>REST</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>1544 ★; 373 commits, latest 2024-12-06</td>
    <td>A RESTful SQLite server.</td>
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
    <td>Apache-2.0</td>
    <td>75198 ★; 30014 commits, latest 2024-12-21</td>
    <td>Uses PostgREST.
<a href="https://supabase.com/docs/guides/self-hosting/docker">Guide to self-Hosting with Docker</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/bradleyboy/tuql">tuql</a></td>
    <td>SQLite 3 or SQL infile</td>
    <td>GraphQL</td>
    <td>JavaScript (Node.js)</td>
    <td>MIT</td>
    <td>1058 ★; 72 commits, latest 2021-06-06</td>
    <td></td>
  </tr>
  <tr>
    <td><a href="https://github.com/nicolasff/webdis">Webdis</a></td>
    <td>Redis</td>
    <td>REST</td>
    <td>C</td>
    <td>BSD-2-Clause</td>
    <td>2861 ★; 667 commits, latest 2024-10-23</td>
    <td>Supports pub/sub with chunked transfer encoding and WebSockets.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/wundergraph/wundergraph">WunderGraph</a></td>
    <td>PostgreSQL, MySQL, SQLite, SQL Server, MongoDB + Atlas, PlanetScale, YugabyteDB,
Neon.tech, FaunaDB</td>
    <td>GraphQL</td>
    <td>TypeScript, Go</td>
    <td>Apache-2.0</td>
    <td>2293 ★; 1920 commits, latest 2024-10-28</td>
    <td><a href="https://docs.wundergraph.com/">WunderGraph Docs</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/xtdb/xtdb">XTDB</a></td>
    <td>XTDB</td>
    <td>REST</td>
    <td>Clojure</td>
    <td>MIT</td>
    <td>2587 ★; 6186 commits, latest 2024-12-19</td>
    <td>A database with a built-in REST API.
<a href="https://hub.docker.com/r/juxt/xtdb-standalone-rocksdb">Official Docker image</a>.</td>
  </tr>
  <tr>
    <td><a href="https://github.com/BjoernKW/ZenQuery">ZenQuery</a></td>
    <td>PostgreSQL, MySQL, IBM Db2, Oracle Database, Microsoft SQL Server and
<a href="https://github.com/BjoernKW/ZenQuery#database">others</a></td>
    <td>REST</td>
    <td>Java</td>
    <td>Apache-2.0</td>
    <td>70 ★; 283 commits, latest 2018-10-16</td>
    <td>Read-only.</td>
  </tr>
</table>


GitHub stats updated 2024-12-21. The commit count and the latest commit date are for the default branch.

# Related projects

For projects that depend on or enhance those on the list see the [Related projects](https://github.com/dbohdan/automatic-api/wiki/Related-projects) wiki page. Feel free to add yours.

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or to update an existing one. See [CONTRIBUTING](./CONTRIBUTING.md) for the details.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
