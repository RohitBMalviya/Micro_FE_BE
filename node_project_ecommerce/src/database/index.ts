import { config } from '@/utils';
import mongoose from 'mongoose';

export default async function connectDB() {
  const connectInstance = await mongoose.connect(
    `${config.MONGODB_URL!}/${config.DATABASE_NAME}`
  );
  try {
    console.log(
      'Mongoose Connected Successfully,',
      connectInstance.connection.host
    );
  } catch (error) {
    console.error('Mongoose Connection Faild Error, ', error);
    process.exit(1);
  }
}
