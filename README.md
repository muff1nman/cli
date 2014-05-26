Piethis.com CLI
===============

Command line tool to interact with the [piethis.com](https://www.piethis.com) API.

This project also contains the Go library you can use in your own projects. Have a look at the [Go Doc](http://godoc.org/github.com/piethis/cli/pie).


Help commands
=============

```
Usage:
  pie [OPTIONS] <command> [command-OPTIONS]

Application Options:
      --db=       The database file to use. (pie.db)
      --url=      The API url prefix, including version. (https://api.piethis.com/v1)
      --raw       Returns raw (json) responses. (false)


Help Options:
  -h, --help      Show this help message

Available commands:
  all-tags
  comments
  company
  login
  my-tags
  new-comment
  new-post
  notifications
  stream

[comments command options]
      -p, --post= ID of the post to add the comment
[company command options]
      -c, --company= ID of the company.
[login command options]
      -e, --email= Your e-mail address to login.
[new-comment command options]
      -t, --text= Text for your new comment
      -p, --post= ID of the post to add the comment
[new-post command options]
      -t, --topic=    Topic to start a new chat.
          --thoughts= First thoughts for the new chat.
```

