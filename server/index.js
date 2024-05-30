require("dotenv").config();
var cors = require("cors");
var express = require("express");
const path = require("path");
var app = express();

const { PORT, TRELLO_TOKEN: TOKEN, TRELLO_API_KEY: API_TOKEN } = process.env;

app.use(cors({ origin: "https://trello.com" }));
app.use(express.static(path.join(__dirname, "public")));
app.use(express.static(path.join(__dirname, "html")));

app.get("/", (req, res) => {
  res.sendFile("index.html");
  // res.send("GET request to the homepage");
});

app.listen(PORT, () => {
  console.log(`Example app listening on PORT ${PORT}`);
});
