The following is a list of tools that automatically expose a REST, GraphQL, or another kind of API for your database.

|                          Project name/link                           |                                       Database(s) supported                                       | API type | Implementation language |                  License                  |                      GitHub stats                       |   Notes    |
|----------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|----------|-------------------------|-------------------------------------------|---------------------------------------------------------|------------|
| [Datasette](https://github.com/simonw/datasette)                     | SQLite 3                                                                                          | REST     | Python 3                | Apache 2.0                                | 1092&nbsp;★; 253&nbsp;commits, latest&nbsp;2017-12-15   | Read-only. |
| [DreamFactory](https://github.com/dreamfactorysoftware/dreamfactory) | MySQL, PostgreSQL, SQLite, MongoDB, CouchDB, and [others](https://www.dreamfactory.com/products). | REST     | PHP 5                   | Apache 2.0, proprietary (optional extras) | 725&nbsp;★; 760&nbsp;commits, latest&nbsp;2017-11-16    |            |
| [PostGraphQL](https://github.com/postgraphql/postgraphql)            | PostgreSQL                                                                                        | GraphQL  | TypeScript (Node.js)    | MIT                                       | 4504&nbsp;★; 665&nbsp;commits, latest&nbsp;2017-12-19   |            |
| [PostgREST](https://github.com/begriffs/postgrest)                   | PostgreSQL                                                                                        | REST     | Haskell                 | MIT                                       | 10067&nbsp;★; 1351&nbsp;commits, latest&nbsp;2017-12-12 |            |
| [sandman2](https://github.com/jeffknupp/sandman2)                    | All supported by SQLAlchemy (MySQL, PostgreSQL, SQLite, Oracle, MS SQL, and others).              | REST     | Python 2/3              | Apache 2.0                                | 644&nbsp;★; 129&nbsp;commits, latest&nbsp;2017-03-06    |            |
| [tuql](https://github.com/bradleyboy/tuql)                           | SQLite 3                                                                                          | GraphQL  | JavaScript (Node.js)    | MIT                                       | 175&nbsp;★; 34&nbsp;commits, latest&nbsp;2017-11-22     | Read-only. |
| [xmysql](https://github.com/o1lab/xmysql)                            | MySQL                                                                                             | REST     | JavaScript (Node.js)    | MIT                                       | 1672&nbsp;★; 205&nbsp;commits, latest&nbsp;2017-12-07   |            |


GitHub stats updated 2017-12-23. The commit count and the latest commit date are for the default branch (usually `master`).

# Contributing

Your contributions are welcome! Please submit a pull request or create an issue to add a new project to the list or update an existing one.

Note that `README.md` is automatically generated from `README.md.template` and the data in [`data/`](./data/). Do not edit `README.md` directly. To update a project's information, edit the corresponding file in `data/`. To add a project, create a new file.

# License

This document and the data in `data/` are licensed under the [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/). By contributing you agree to release your contribution under this license.
