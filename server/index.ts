import express, { Request, Response } from 'express';
const dotenv = require('dotenv');

dotenv.config();

const app = express();
const port = process.env.MQTT_HTTP_PORT;
var req_count = 0;

app.get('/', (req : Request, res : Response) => {
  req_count++;
  console.log(req_count);
  res.send(req_count.toString());
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
