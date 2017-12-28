# How to edit the list

`README.md` is automatically generated from `README.md.template` and the data in [`data/`](./data/). Do not edit `README.md` directly. To update a project's information, edit the corresponding file in `data/`. To add a project, create a new file.

# What projects to submit

(Below, "must" means a hard requirement and "should" means a preference.)

To qualify as "automatic" for our purposes, a piece of software that exposes the API must take a minimal amount of configuration to set up. A good example would be just the credentials to access the database and a port on which to serve the API. If the database itself has a schema (like SQL databases do), it must not require the user to define schemas or models in the configuration to get started and should not require the user to define the API endpoints. It should respond to changes to the database schema with no or minimal user intervention (e.g., a restart or a SIGHUP). This means that most Web frameworks, code generators, and boilerplates don't qualify.

A submitted project must be at least alpha quality, that is, ready for users to test. A project that is no longer actively maintained must be more mature than that and already tested by users.

Please only submit open source or [open core](https://en.wikipedia.org/wiki/Open_core) software you can run on your own servers.
