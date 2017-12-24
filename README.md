The following is a list of tools that automatically expose a REST, GraphQL, or another kind of API for your database.

|                          Project name/link                           |                                       Database(s) supported                                       | API type | Implementation language |                  License                  |                      GitHub stats                       |   Notes    |
|----------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|----------|-------------------------|-------------------------------------------|---------------------------------------------------------|------------|
| [Datasette](https://github.com/simonw/datasette)                     | SQLite 3                                                                                          | REST     | Python 3                | Apache 2.0                                | 1097&nbsp;★; 253&nbsp;commits, latest&nbsp;2017-12-15   | Read-only. |
| [DreamFactory](https://github.com/dreamfactorysoftware/dreamfactory) | MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and [others](https://www.dreamfactory.com/products). | REST     | PHP 5                   | Apache 2.0, proprietary (optional extras) | 727&nbsp;★; 760&nbsp;commits, latest&nbsp;2017-11-16    |            |
| [PHP-CRUD-API](https://github.com/mevdschee/php-crud-api)            | MySQL, PostgreSQL, MS SQL Server 2012. Limited support for SQLite 3.                              | REST     | PHP 5                   | MIT                                       | 1226&nbsp;★; 1004&nbsp;commits, latest&nbsp;2017-11-19  |            |
| [PostGraphQL](https://github.com/postgraphql/postgraphql)            | PostgreSQL                                                                                        | GraphQL  | TypeScript (Node.js)    | MIT                                       | 4510&nbsp;★; 665&nbsp;commits, latest&nbsp;2017-12-19   |            |
| [PostgREST](https://github.com/begriffs/postgrest)                   | PostgreSQL                                                                                        | REST     | Haskell                 | MIT                                       | 10074&nbsp;★; 1351&nbsp;commits, latest&nbsp;2017-12-12 |            |
| [pREST](https://github.com/prest/prest)                              | PostgreSQL                                                                                        | REST     | Go                      | MIT                                       | 1560&nbsp;★; 400&nbsp;commits, latest&nbsp;2017-12-20   |            |
| [sandman2](https://github.com/jeffknupp/sandman2)                    | All supported by SQLAlchemy (MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).              | REST     | Python 2/3              | Apache 2.0                                | 645&nbsp;★; 129&nbsp;commits, latest&nbsp;2017-03-06    |            |
| [tuql](https://github.com/bradleyboy/tuql)                           | SQLite 3                                                                                          | GraphQL  | JavaScript (Node.js)    | MIT                                       | 176&nbsp;★; 34&nbsp;commits, latest&nbsp;2017-11-22     | Read-only. |
| [xmysql](https://github.com/o1lab/xmysql)                            | MySQL                                                                                             | REST     | JavaScript (Node.js)    | MIT                                       | 1678&nbsp;★; 205&nbsp;commits, latest&nbsp;2017-12-07   |            |


GitHub stats updated 2017-12-24. The commit count and the latest commit date are for the default branch (usually `master`).

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or to update an existing one.

Note that `README.md` is automatically generated from `README.md.template` and the data in [`data/`](./data/). Do not edit `README.md` directly. To update a project's information, edit the corresponding file in `data/`. To add a project, create a new file.

## What projects to submit

To qualify as "automatic" for our purposes, a piece of software must be a server that requires a small amount of configuration to set up, not a Web framework for building APIs, a code generator that leaves maintaining the code to its user, or a boilerplate. It should respond to changes to the database schema with no or minimal user intervention (like a restart or a SIGHUP). A submitted project must be at least alpha quality, i.e., ready for users to test. A project that is no longer actively maintained must be more mature than that and already tested by users.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
