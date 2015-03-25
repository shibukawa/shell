Shell Style Text Processing
==================================

Shell-ish text processing algorithm collection

* Splitter that supports quoted string.
* Escape/Unescape string for shell

Install
----------

.. code-block:: bash

   $ go get github.com/shibukawa/shell

Reference
--------------

Split
~~~~~~~~~

``func Parse(src string) []string``

Split src string by space. If the chunk is quoted, the spaces between double-quotations are ignored.

.. code-block:: go

   import "github.com/shibukawa/shell"

   func parse() {
       opts := shell.Parse(`foo bar baz "hello world"`)
       // opts[0] = "foo"
       // opts[1] = "bar"
       // opts[2] = "baz"
       // opts[3] = "hello world"
   }

Escape
~~~~~~~~~

``func Escape(src string) string``

Escape src string.

If src string contains the following characters, it quotes the src and escape some characters:

* ``"`` ``$`` ``@`` ``&`` ``'`` ``(`` ``)`` ``^`` ``|`` ``[`` ``]`` ``{`` ``}`` ``;`` ``*`` ``?`` ``<`` ``>`` ``\``` ``\``
* Escape sequences: ``\r`` ``\n`` ``\t``
* White space

Unescape
~~~~~~~~~

``func Unescape(quotedSrc string) string``

It return unescaped string.

License
------------

MIT
