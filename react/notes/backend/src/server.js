import express from 'express';
import dotenv from 'dotenv';
import cors from 'cors';

import notesRoutes from './routes/notesRoutes.js';
import { connectDB } from './config/db.js';
import rateLimiter from './middleware/rateLimiter.js';

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;


// Middleware
app.use(cors(
  {
    origin: process.env.CORS_ORIGIN || 'http://localhost:5173',
  }
));
app.use(express.json());
app.use(rateLimiter);

app.use("/api/notes", notesRoutes)

connectDB().then(() => {
  app.listen(PORT, () => {
    console.log("Server is running on http://localhost:"+PORT);
  });
});