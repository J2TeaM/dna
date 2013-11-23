/*
 This package is supposed to implements wrapper for postgresql manipulation.
 It intends to leave the standard package alone. It inherits from sql.Rows or sql.Db.

 Name convention once converting from names of a struct to columns' names in a table:

 	* Column's name is always be snake case. Ex: date_created
 	* Struct field's name which is exported always starts with upper camel case. Ex: DateCreated
*/
package sqlpg
