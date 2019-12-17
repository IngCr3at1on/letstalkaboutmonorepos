# Let's Talk about Mono Repos

They suck; seriously, anyone who tries to convince you otherwise has loads of tooling designed to make their day better but in general they're really not as good of an idea as most people will try to convince you.

## What is this repo?

The purpose of this repo is to house documentation resources (and partial examples) for breaking a monorepo of multiple services into seperate single service repos. This is a method I've used in production to pull services out of a monorepo containing 40+ services into their own standalone repos for service deployment on kubernetes.

While the code included will be in Go this is just to include examples to show that what we've done works, the same process of breaking down a monorepo should apply for other languages.

## What knowledge should I have before attempting to follow these steps?

Don't be afraid to use a command line interface, we will be using it for every single step discussed.

Know git basics; you don't need to know how to filter a repo down (that's the point of this) but you should probably understand basic commit/push practices.
