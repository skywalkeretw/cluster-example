const express = require('express')

const app = express()
const port = 8080
app.use(express.urlencoded({extended: true}));
app.use(express.json())

app.get('/', (req, res) => {
    res.setHeader('Content-Type', 'application/json');
    res.status(200).send(JSON.stringify({"body": "Hello from Node"}))
})



app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})
