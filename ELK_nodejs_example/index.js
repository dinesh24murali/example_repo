var express = require("express");
var logger = require("./logger");
var app = express();

app.get("/", function (req, res) {
  logger.info('Inside hello world');
  res.send("Hello world!");
});

app.get("/error_or_not/:text", function (req, res) {
  logger.info('Inside error_or_not');
  logger.info(`Received text: ${req.params.text}`);
  if (req.params.text === "error") {
    logger.error('Received error Request');
    res.status(500).send(req.params.text);
    return;
  }
  logger.info('Received normal Request');
  res.status(200).send(req.params.text);
});

app.listen(3000);
