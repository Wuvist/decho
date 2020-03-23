# decho

`decho` is a demo project utilizing:

* Dependency Injection with [wire](https://github.com/google/wire)
* Web framework [echo](https://echo.labstack.com).
* Configuration with [toml](https://github.com/BurntSushi/toml)
* DB ORM with [sqlboiler](https://github.com/volatiletech/sqlboiler)
* HMLT Template engine with [gorazor](https://github.com/sipin/gorazor)

The idea is to build modularized & testable web applicaiton, and it's inspired by [micronaut](https://micronaut.io/).

# Todo

* [X] Add sql db to controller
  * [ ] Add mock
* [ ] Add test case
* [X] How to add multiple controller?
* [X] provider set for controller
* [ ] How to make config loading as an lib?
* [X] refactor newWebApp with field injection
