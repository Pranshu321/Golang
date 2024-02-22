const { json } = require("body-parser");
const express = require("express");
const app = express();
const bodyParser = require("body-parser");

app.use(bodyParser.urlencoded());
app.use(bodyParser.json());
app.use(express.json())

app.get("/", (req, res) => {
    res.json({
        "status": "Server Working"
    })
})

app.get("/get", (req, res) => {
    res.json({
        message: "Hello World"
    });
});

app.post("/post", (req, res) => {
    console.log(req.body);
    const body = req.body;
    res.json(body)
});

app.post("/postform" , (req, res) => {
    res.json(req.body);
})

app.listen(3000);