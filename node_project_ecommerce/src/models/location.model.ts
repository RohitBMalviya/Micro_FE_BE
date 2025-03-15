import mongoose from 'mongoose';

const locationSchema = new mongoose.Schema(
  {
    label: {
      type: String,
      unique: true,
      select: true,
    },
    parend_id: {
      type: mongoose.Types.ObjectId,
      unique: true,
      select: true,
    },
    deleted_at: {
      type: Date,
      select: false,
      default: null,
    },
  },
  { timestamps: true }
);

const Location = new mongoose.Model('Location', locationSchema);

export default Location;
