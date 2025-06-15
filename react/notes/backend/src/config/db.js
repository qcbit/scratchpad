import mongoose from 'mongoose';

export const connectDB = async () => {
  try {
    await mongoose.connect('mongodb+srv://'+process.env.MONGO_USER+':'+process.env.MONGO_PASSWORD+'@'+process.env.MONGO_URI || 'MONGO_URI_NOT_SET');
    console.log("MongoDB connected successfully");
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
};