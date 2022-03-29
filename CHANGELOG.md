TO-DO
=====
- [X] Add single `ctrl +c` to stop both.


FEATURE 15/11/2021
==================
- Added context to handle management of both cert and token renewal threads


FEATURE 15/11/2021
==================
- Seperate function for certificate management as certs cannot be renewed and need to reissued.
- buffered channels for signal notify

BUG-FIX 15/11/2021
==================
- Had added token into both watchers