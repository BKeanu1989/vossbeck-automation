require("dotenv").config();
var cors = require("cors");
var express = require("express");
var app = express();

const { PORT, TRELLO_TOKEN: TOKEN, TRELLO_API_KEY: API_TOKEN } = process.env;

app.use(cors({ origin: "https://trello.com" }));

console.log(PORT, TOKEN);
app.get("/", (req, res) => {
  res.send("GET request to the homepage");
});

app.listen(PORT, () => {
  console.log(`Example app listening on PORT ${PORT}`);
});
