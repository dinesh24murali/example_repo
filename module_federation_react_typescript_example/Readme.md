# Module federation:

Project was setup taking reference from this [blog](https://medium.com/walmartglobaltech/module-federation-using-webpack-5-the-micro-frontend-journey-719688c5d73b)

These are the github repos that the article points to

Left Nav micro-frontend application: https://github.com/priyavarun/wp5-mf-left-nav

Top Nav micro-frontend application: https://github.com/priyavarun/wp5-mf-top-nav

Item details micro-frontend application: https://github.com/priyavarun/wp5-mf-item-details

Shell container: https://github.com/priyavarun/wp5-mf-shell

# How to run?

Step 1:

Open `left-nav`, `right-nav`, and `shell` in individual terminals.

Step 2:

Use node version mentioned in `.nvmrc`

Step 3:

Do `npm i` in all three folders

Step 4:

`Shell` is the project that hosts `left-nav`, and `right-nav`

Run `npm start` in all the three terminals

open http://localhost:3004 to access the `shell` app
