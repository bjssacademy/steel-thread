# Using this code

If you have cloned this repository to your local machine to use, then to be able to use it, there are a few things you need to do.

The code provided is 2 separate projects, and as such the *root* for each is inside the code-postgres or code-frontend folders.

1. In the `code-postgres` folder run the command `git init` to initialize a new git repository.

2. Create a new repository on github/wherever you are pushing the code to, ensuring it is empty.

3. Run the following commands:

```bash
git add .
git commit -m "initial"
```

4. Now add the remote:

```bash
git remote add origin your.repository.url
```

5. Now we have a remote, we can push our code to it using `git push --set-upstream origin master`.

## Add .env file

Create a new `.env` file and fill out the connection details. 

```bash
DBTYPE="postgres"
DBHOST="YOURDBHOSTNAME.postgres.database.azure.com"
DBNAME="YOURDBNAME"
DBUSER="YOURDBUSER"
DBPASSWORD="YOURDBPASSWORD"
DBSSLMODE="require"
```

> :exclamation: You have to do this as the `.env` file is listed in the `.gitignore` file so your secrets don't get pushed to you repository!

## Add the azure-pipelines.yml

Create the file `azure-pipeline.yml`; you can now continue with the instructions in [steel threading](../readme.md#connect-to-our-remote-db).
