# Friendbot Service for the AiBlocks Test Network

Friendbot helps users of the AiBlocks testnet by exposing a REST endpoint that creates & funds new accounts.

Millennium needs to be started with the following command line param: --friendbot-url="http://localhost:8004/"
This will forward any query params received against /friendbot to the friendbot instance.
The ideal setup for millennium is to proxy all requests to the /friendbot url to the friendbot service
