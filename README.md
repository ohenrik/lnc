# LNC - LN.Capital

**⚠️ This is alpha software and should not yet be used in production ⚠️**

LNC is a lightning network capital management tool built for routing nodes.
The ultimate goal is to create a set of tools that combine node & channel
introspection, capital management and network analysis accross a set of tools 
that simplifies the task of operating a routing node. This goal might change, 
based on input from the community.

If you want to contribute to this project, feel free to join the 
[Telegram group](https://t.me/joinchat/V-Dks6zjBK4xZWY0), fork this repo or
submit a feature request.

## Roadmap

This is a hight level overview of possible future features. Feel free to 
suggest new features.

* HTLC stream and store to database from LND
* Combine transaction, channel, fee and htlc data
* Create profit and loss overview.
* Channel categorization and grouping
* Fee logic/ruless
* Rebalancing logic/rules
* A rest API
* Command line interface
* Visual interface (Webinterface)
* Add support for c-lightning

## Build

````bash
git clone https://github.com/lncapital/lnc.git
cd lnc
go build /cmd/lnc/lnc.go
````

