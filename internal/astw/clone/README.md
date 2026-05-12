# astw/copy

-   Concrete copiers creates new instances of same type by indirect the pointer value. Later; calls special copiers for copying the slice and map type fields.

-   Interface copiers further calls concrete copiers after type check on src.

-   Slice and map copiers are used by concrete copiers as explained above.
