var express = require("express");
var app = express();

app.get("/", function (req, res) {
  res.send("Hello world!");
});

app.get("/error_or_not/:text", function (req, res) {
  if (req.params.text === "error") {
    res.status(500).send(req.params.text);
    return;
  }
  res.status(200).send(req.params.text);
});

app.listen(3000);
