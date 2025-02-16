import { config } from './utils';
import app from './app';
import dotenv from 'dotenv';
import connectDB from './database';

dotenv.config({ path: './env' });

const port = config.PORT || 4001;

connectDB()
  .then(() => console.log('Database Connected Successfully !!!.'))
  .catch((error) => console.log('Database Connection Failed,', error));

app.listen(port, () => {
  console.log(`Server Is Running On Port http://localhost:${port}`);
});
