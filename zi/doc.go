/*
This package implements some functionalities regarding mp3.zing.vn.

It downloads metadata from the site using 2 methods:
Getting content from main site by filtering HTML tags or from API. The API is default setting.

It also encrypts, decrypts id, key or encoded key.
Due to some mistakes in current API, it is able to find durect links of songs, videos or tv video clips.
But the exploitation cannot remain for long time. So it has to be checked again.

Basic types of this package are: Song, Video, Artist, Album, TV that conforms site.Item interface.
All types with prefix "API" mean the are the types representing items from API and have to be converted to the basic types of the package.

The difference between basic types' ids and real ids from the site is a constant : ID_DIFFERENCE =  307843200

PROGRESS: type TV, Artist are not complete!!!!!
*/
package zi
