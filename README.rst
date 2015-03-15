OptString Parser
======================

Shell-ish string splitter that supports quoted string.

Install
----------

.. code-block:: bash

   $ go get github.com/shibukawa/optstring_parser

Usage
---------

.. code-block:: go

   import "github.com/shibukawa/optstring_parser"

   func parse() {
       opts := optstring_parser.Parse(`foo bar baz "hello world"`)
       // opts[0] = "foo"
       // opts[1] = "bar"
       // opts[2] = "baz"
       // opts[3] = "hello world"
   }

Reference
--------------

``func Parse(src string) []string``

Splits src string by space. If the chunk is quoted, the spaces between double-quotations are ignored.

License
------------

MIT
