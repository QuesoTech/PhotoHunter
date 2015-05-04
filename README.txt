Environment
===========
All work was done in a linux command line environment.

Photohunter
===========
Build
-----
Photohunter requires:
- NodeJS
- Apache Cordova
- Android SDK
- Android Version 4.4.2
- phonegap-facebook-plugin

To deploy straight to an Android (4.4.2) phone, make sure the phone is connected
and is in debug mode. Then, open a terminal and execute the following commands:
> cd path/to/project/folder/PhotoHunter/src/photohunter
> cordova run android

To just build the apk, run the following commands:
> cd path/to/project/folder/PhotoHunter/src/photohunter
> cordova build
The build output will display a path to the built apk.

Installation
------------
The PhotoHunter application can be installed by downloading the apk file to an
Android device (Version 4.4.2) and installing it through the phone's interface.

Use
---
Upon launch, the app will display the title screen with a log in button.
Clicking the login button will prompt the user to login using their Facebook
credentials.  Once logged in, the user is presented with a list of categories
to choose from, as well as a button (their Facebook profile picture) linking to
their account page.  Clicking on the profile picture button sends the user to a
page with their account information, filled out using the Facebook API.
Clicking on a category in the category list sends the user to a photo capture
page.  Clicking the capture button opens the native camera application on the
user's phone and allows the user to take a picture.  Once the picture is taken,
an upload button and the user's location appear.  Clicking on the upload button
sends the image to the database.

QuickPic
===========
Build
-----
QuickPic requires:
- NodeJS
- Gulp (npm install -g gulp-cli)
- Apache Cordova
- Android SDK
- Android Version 4.4.2
- phonegap-facebook-plugin

To deploy straight to an Android (4.4.2) phone, make sure the phone is connected
and is in debug mode. Then, open a terminal and execute the following commands:
> cd path/to/project/folder/PhotoHunter/src/quickpic
> npm install # once only
> gulp deploy-local

To just build the apk, run the following commands:
> cd path/to/project/folder/PhotoHunter/src/quickpic
> npm install # once only
> gulp emulate
The build output will display a path to the built apk. You can close the
emulator that is spawned.

Installation
------------
The QuickPic application can be installed by downloading the apk file to an
Android device (Version 4.4.2) and installing it through the phone's
interface. Upon launch, the app will display the title screen with a log in
button. Clicking the log in button will prompt the user to log in with their
Facebook account.

Use
---
Once logged in, the user is presented with the main menu. Three buttons are
displayed: Photo Quiz, Statistics, and Leaderboards. However, only the 'Photo
Quiz' button is implemented at this time. Clicking this button will begin a
photo quiz. An image will be displayed for a very short period (presently
750ms), and then a list of potential categories will be presented to the
user. Upon selecting a category, statistics about the distribution of category
choices are presented. This statistics screen has two buttons: Menu and
Continue. Menu returns to the main menu, and Continue presents the user with
another image.


=======



Researcher's Interface
======================
The Researcher's Interface requires the following package be already installed:
- Golang, version 1.2 or greater
- PostgresSQL, 9.4.1
- PostGIS, 2.0.7
These package are available as source code on their respective project pages,
and as binaries that can be installed from all major Linux package managers.
Their installation is outside the purview of this document.

Installation
------------
The Researcher's Interface can be installed by downloading the source code from
http://www.github.com/QuesoTech/PhotoHunter.  Then performing the following
commands in a Bash shell:
> cd src/researcherinterface
> go build
> ./researcherinterface

The server will now be running, listening to port 8080. There are a number of common
methods for routing http requests to this port.

Use
---
Upon arrival to the sites landing page, a user may either create a new account or
log in to an existing one. On signed in, a dashboard will appear. From here
a reasercher may fill out a form requesting a new dataset, manage existing datasets,
download completed datasets, or sign out.
