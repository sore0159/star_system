## 2017 09 26

### What progress has been made?

I have built a 'manage' tool to aid in the creation and adjustment of postgres tables and data.  I've started a basic Star data creation routine and tested loading/reading from the database.

### What interesting challenges have been encountered?

* Getting postgres operational and synchronized between development and 'production' environments was a challenge, though not really part of this codebase.

* Creating a separate binary for management of tables and star data came about after considering separation of concerns; there's no reason these tasks need to be part of any data serving/authentication.  

* When creating the management program I had to consider the appropriate level of abstraction: in the past I have gone a bit overboard in trying to abstract away database code, creating interfaces for managing databases for semi-arbitrary go data structures.  My old [database utility package](https://github.com/sore0159/golang-helpers) was a fundamental part of my previous large web app design, but this time I felt a more concrete design would be better.

* I almost wrote another, smaller query generator, but decided to go with manual table schema for the time being.  I intend to use manual queries for table usage, as well.

* I looked at [pgx](https://github.com/jackc/pgx) for database driver use, but decided on the more basic standard library postgres sql driver.  I am, however, trying to create a slight wrapper around the db to hopefully make future changes at this level less painful (though I am not optimistic).

* Location coordinates have switched from uint64 to int64 to better use postgresql's bigint data type.  I'm uncertain about the captain UID data type: bigserial is the most straightforward choice on the db level, but this is essentially just an int64, and server side I have already built around the theoretically unbounded BigInt.  I will try a 'numeric' database data type for the UID, with concerns about efficiency.

* I had to consider how to go about getting the star data from generating program into the database.  The intended scale of this data is such that holding it in memory is not an option, so I am currently generating data in small chunks and then writing them to a file, which upon completion is fed to Postgres' ``COPY FROM`` command.  I'm trusting postgresql to know how to safely read data from a huge file: this has not yet been tested at scale.

    ``COPY FROM`` has some problems already: primarily requiring a superuser database role to allow postgresql to read from the filesystem.  Combined with difficulty in verifying uniqueness in chunk-generated data, it I might investigate how slow ``INSERT`` statements actually are when I start generating data sets large enough.

## 2017 09 07

### What progress has been made?

"Captain" data structures have been added using the std's big.Int type for UID, with a mock DB structure and a basic secure-cookie authentication system.  The test server recognizes first and repeat visits.  JSON loading/saving of the mock db works.

The test server has been structured to avoid the use of global variables as much as possible.  Sub-packages, like the current 'captains' package, take as function parameters things like loggers, secure keys, and data files. 

The test server has a 'configurations' data structure that at present simply uses a list of default values, but is made to be expandable into command line or rc-file configuration.

### What interesting challenges have been encountered?

1. The shift away from storing the logger and the secure cookie encoder/decoder in global variables occurred because of the separation of design in the package system.  The captains package started out relying on an assumed shared structure with whatever server would be running it, but when building the Route Builder I had to consider how to handle logging.

If the route builder should have access to the (then) global logger, then it should be a part of the server package itself.  I considered splitting off logging/routing from the test server, but decided to instead keep the captains package as separate as possible by having the route builder take a logger (interface) as a parameter.

This done, transforming the logger from a global to a server-held variable seemed a natural next step, along with the rest of the program resources.

2. Another interesting consideration was the UID type for the captains data structure.  Given the semi-disposable planned nature for user identity, I decided that an unbounded ID count was worth the extra work in using the big.Int type instead of a simple uint64.  The intention is to store this in a postgres bigserial type.

3. Finally, the encryption/security of the authentication system is being handled by the [gorilla/securecookie](http://www.gorillatoolkit.org/pkg/securecookie) package.  I know not to 'roll my own crypto', though I'm not sure how much I want to integrate toolkits like gorilla into my program.  For now I'm using it as simply as possible, for only the actual encrypting of the UID values.

---
## 2017 09 06

### What is the project?

Project "Star System" (working title) is a long-term project to build and deploy a web-app allowing the public to cooperatively explore a large star system.  Inspirations: No Man's Sky, [Space Engine](http://spaceengine.org/), 4X Space Games.

### What are the design goals of the project?

* Demonstrate/exercise knowledge of web-app design and deployment in a general sense.
* Learn/practice 3D webGL UI development.
* Build something that my friends and I enjoy playing with.
* Learn/practice large scale (relatively) database usage.

### What phases of development do I anticipate?

* Authentication: develop a system of password-less disposable identities backed by secure cookies.
* Persistence: Set up a Postgres database for the project, ready to scale to the needs of the project as it develops.
* Data: Populate the project with simple star data, expand the system to user-specific data such as "location" and "destination".
* UI: Develop a webGL display of the data.
* Interaction: Develop a system of user commands; build the app system to implement them and expand the UI to accommodate them.
* Deployment: Find an external host.  Modify the system to fit the needs of the host.  Set up the server app, database, and https on this host.
* User testing: Get some people to use it, and get data from their usage.

### Challenges:

* Self motivation/focus
* Leaning WebGL
* Scaling data to a meaningfully large size
* Working on an externally hosted system
* Creating a useful display of a star field and associated data
* Making the product fun to use



