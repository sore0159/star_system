### What is the project?

Project "Star System" (working title) is a long-term project to build and deploy a web-app allowing the public to cooperatively explore a large star system.  Inspirations: No Man's Sky, [Space Engine](http://spaceengine.org/), 4X Space Games.

### What can it do right now?

Running go build in the star\_system/test\_server directory will create a binary that, when run in the same directory as the included FILES folder, will serve a basic page that will set/read a secure cookie identifying each visitor with a 'captain'.  Log files and data files will be generated in the FILES folder.

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



