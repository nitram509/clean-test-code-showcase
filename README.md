# golang-clean-test-code-showcase

A showcase, of how unit test can be written.
No claim to be complete here.

### Business requirement - for this demo purpose

This showcase/demo library has some acceptance criteria

* as a developer I would like to extract meaningful information from a condensed firmware string literal
* the information extraction must work on a single line string
* the information extraction must work on multi-line strings, separated by newline (\\n) 
* these fields must be available in the meaningful extracted information (struct)
   * FirmwareName,TargetName,TargetDetail,Version,ReleaseDateStr,GitHash,ReleaseTime

Example firmware String: \
```Betaflight / SPRACINGF3EVO (SPEV) 3.4.0 Apr 17 2018 / 14:00:13 (b2c247d34) MSP API: 1.39```

### Implementations

1. Using native/vanilla GO language and library
2. Using famous [stretchr/testify](https://github.com/stretchr/testify) library
3. Using broken and outdated [rdrdr/hamcrest](https://github.com/corbym/gocrest) library
4. Using almost unknown [corbym/gocrest](https://github.com/corbym/gocrest) library

### What is / Why "HAMCREST"?

[Hamcrest](http://hamcrest.org/) is a famous Java Library, that was developed to make
unit tests in Java more readable and follow [Clean Code principles](https://de.wikipedia.org/wiki/Clean_Code)

Cite from hamcrest.org ... \
*Matchers that can be combined to create flexible expressions of intent* \
*Born in Java, Hamcrest now has implementations in a number of languages*
