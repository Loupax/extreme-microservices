# extreme-microservices
## (Or what is the maximum level of granularity I can reach when building a web application)

I catch myself repeating the same matra. The only meaningfully reusable unit of software is the binary itself. What if I put this idea to the test by doing the most extreme implementation possible?

Apparently I need to do some reading on the purpose of microservices, domain driven design and MVC in order to write a good preamble for that project so I will be efficient and won't bother for now. 

Long story short, I want to see how granular I can get when developing a complex backend. I want to find what will happen if I push the microservice paradigm to it's extreme and see what I design ideas I can move from this silly pet project to a professional environment.

First things first, every endpoint API must be handled by a separate binary, and should have as little knowledge outside it's domain as possible.

Ideally, all routes should be defined on the infrastructure level, let is be by kubernetes ingress rules or nginx proxy rules.

### Rules of the game
* Each endpoint should be it's own binary. 
* Each endpoint should be able to run with or without a path in it's URL
* Before adding exceptions to existing binaries, see if we can push that handling in the infrastructure level; i.e. by creating a separate binary and have a proxy or the cluster decide where to forward our request
* Even though this application expects to be run on a cluster, development should be easy and require minimal setup.
* Even though each endpoint is it's own binary, we should be able to deploy the application as a single container if we wish to
* We should try our best to not share data between endpoints by accessing database directly. Each application should still be isolated from the rest, even though they are part of the same repository.

# So, what are we making?
This monorepo will implement the backend of a turned based strategy game that mimics the mechanics of [Advance Wars](https://en.wikipedia.org/wiki/Advance_Wars). Why? Because I like the game. 

API clients will be able to:

* Create game maps with multiple possible kinds of terrain, buildings and units of varying speeds, attack types an defense
* Start battles between armies on those maps
* Move units on the map and attack
