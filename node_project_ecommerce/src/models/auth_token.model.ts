import mongoose from 'mongoose';

const authTokenSchema = new mongoose.Schema(
  {
    user_id: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'User',
    },
    refresh_token: {
      type: String,
      select: true,
      unique: true,
    },
    refresh_token_expiry: {
      type: Date,
      select: false,
    },
    deleted_at: {
      type: Date,
      select: false,
      default: null,
    },
  },
  { timestamps: true }
);

const AuthToken = new mongoose.Model('AuthToken', authTokenSchema);

export default AuthToken;
