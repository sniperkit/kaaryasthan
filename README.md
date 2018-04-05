# Kaaryasthan - Task Management for Small Teams
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaaryasthan/kaaryasthan)](https://goreportcard.com/report/github.com/kaaryasthan/kaaryasthan)
[![Build Status](https://travis-ci.org/kaaryasthan/kaaryasthan.svg?branch=master)](https://travis-ci.org/kaaryasthan/kaaryasthan)
[![Open Source Helpers](https://www.codetriage.com/kaaryasthan/kaaryasthan/badges/users.svg)](https://www.codetriage.com/kaaryasthan/kaaryasthan)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/kaaryasthan/kaaryasthan/blob/master/CONTRIBUTING.md)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Kaaryasthan helps you to manage private projects.  You can add tasks &
issues (items) in to your project.  An item has title, description,
discussions, creator, assignees, and labels.  Project milestones can
be created with due date and items with priority order.

This repository contains the source code of Kaaryasthan.  It has
source code for both user interface and server.  The user interface of
Kaaryasthan is written in [Angular], server in [Go], and [PostgreSQL]
is used for the database.

> Kaaryasthan (കാര്യസ്ഥൻ) is a [Malayalam] word with meaning "manager".

## Key Features

- All URLs requires authentication except login & registration.
  (Note: This makes it unsuitable for open source projects with public
  issues.  Only private projects can be hosted using Kaaryasthan)
- Threaded discussion. Multiple discussions can be added under an
  item.  And comments can be added under each discussion.
- Easy deployment.  Entire application including front-end is
  available as a single binary.  You also need to install PostgreSQL
  and NGINX unless you are using hosted PostgreSQL service.

## Development

This project requires [Go] version 1.8 or above.  This project also
requires [Node] version 6.11 or above, preferably an LTS release.

Once Go and Node is installed, you can install these utilities:

- <https://github.com/pilu/fresh>
- <https://github.com/jteeuwen/go-bindata>
- <https://github.com/elazarl/go-bindata-assetfs>
- <https://github.com/alecthomas/gometalinter>
- <https://glide.sh>
- <https://cli.angular.io>

To install the above utilities:

    curl https://glide.sh/get | sh
    go get github.com/pilu/fresh
    go get github.com/jteeuwen/go-bindata/...
    go get github.com/elazarl/go-bindata-assetfs/...
    go get github.com/alecthomas/gometalinter
    gometalinter --install
    npm install -g @angular/cli

You can clone [Kaaryasthan] repository inside `$GOPATH` using these
commands (Note: `$GOPATH` should be pointing to a single directory):

    mkdir -p $GOPATH/src/github.com/kaaryasthan
    cd $GOPATH/src/github.com/kaaryasthan
    git clone https://github.com/kaaryasthan/kaaryasthan.git

Now you can run `./build.sh` command.

    cd $GOPATH/src/github.com/kaaryasthan/kaaryasthan
    ./build.sh

Install [Docker] and [Docker Compose] and then run:

    docker-compose up -d

Migrate database schema by running:

    ./kaaryasthan -migrate

To run test cases:

    ./runtests.sh

### Running development servers

To run the back-end development server (from the top-level directory):

    fresh

The back-end will listen on [8080] port.  Whenever there is a change
in Go source files, the server will be automatically restarted.

To run the front-end development server (from the `./web` directory):

    npm start

The web user interface will be available on [4200] port.  You can use
Firefox or Chrome to open it.  Any change in source files will refresh
the user interface automatically.  There is a webhook which does this
magic.

## Contributors

Thanks goes to these wonderful people ([emoji key][emojis]):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars3.githubusercontent.com/u/121129?v=4" width="100px;"/><br /><sub><b>Baiju Muthukadan</b></sub>](http://muthukadan.net)<br />[📖](https://github.com/baijum/kaaryasthan/commits?author=baijum "Documentation") |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors][all-contributors] specification.
Contributions of any kind are welcome!

If you are looking forward to contribute to this project, please take
a look at the [CONTRIBUTING.md] and [wiki page about contributing] for
more deatils.

## FAQ

### Why this project?

I started this as a [pet project] for learning web application development
using Go & Angular.  In fact, I had started this [project in 2014] with
another name.

### Are you seeking external contributions?

Yes! You are welcome to contribute :-)

Please take a look at the [CONTRIBUTING.md] and
[wiki page about contributing] for more deatils.

### Can I use Kaaryasthan to manage open source project tasks & issues?

I think that may not be possible.  Because all URLs will require
authentication except login & registration.  Since open source
projects requires public trackers, this won't be a desirable solution.
Kaaryasthan is designed exclusively for private projects.  And I am
not planning to change it. These are few reasons for not supporting
open source projects:

1. There are many good trackers available for open source projects.
2. Kaaryasthan is not exclusively designed for software projects.
3. Don't want to make the system complex to handle Slashdot effects.
4. No plan for search engine optimization (SEO for Google, Bing etc.).
5. Reduce the scope to keep the software as simple as possible.

The architecture doesn't support Kaaryasthan being used as a public
tracker. I explained this much in-order to not receive any feature
request for the same :-)

---

> [kaaryasthan.org](https://kaaryasthan.org) &nbsp;&middot;&nbsp;
> Demo [demo.kaaryasthan.org](https://demo.kaaryasthan.org) &nbsp;&middot;&nbsp;
> IRC [#kaaryasthan@freenode](https://riot.im/app/#/room/#freenode_#kaaryasthan:matrix.org) &nbsp;&middot;&nbsp;
> [Mailing List](https://groups.google.com/forum/#!forum/kaaryasthan) &nbsp;&middot;&nbsp;
> Twitter [@kaaryasthan](https://twitter.com/kaaryasthan)

[pet project]: https://team-coder.com/pet-project
[Node]: https://nodejs.org/en
[Angular]: https://angular.io
[Go]: https://golang.org
[PostgreSQL]: https://www.postgresql.org
[Malayalam]: https://en.wikipedia.org/wiki/Malayalam
[Docker]: https://docs.docker.com
[Docker Compose]: https://docs.docker.com/compose
[Kaaryasthan]: https://github.com/kaaryasthan/kaaryasthan
[project in 2014]: https://github.com/baijum/pitracker
[CONTRIBUTING.md]: https://github.com/kaaryasthan/kaaryasthan/blob/master/CONTRIBUTING.md
[wiki page about contributing]: https://github.com/kaaryasthan/kaaryasthan/wiki/Contributing
[emojis]: https://github.com/kentcdodds/all-contributors#emoji-key
[all-contributors]: https://github.com/kentcdodds/all-contributors
[8080]: http://localhost:8080
[4200]: http://localhost:4200
