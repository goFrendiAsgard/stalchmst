const express = require("express");
const app = express();

const port = process.env.HTTP_PORT || 3000;

app.get("/", function (req, res) {
  res.send("Hello World!");
});

app.listen(port, function () {
  console.log(`API run at ${port}`);
});
