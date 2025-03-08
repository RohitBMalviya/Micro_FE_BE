import mongoose from 'mongoose';

const userSchema = new mongoose.Schema(
  {
    email: {
      type: String,
      required: true,
      trim: true,
      select: true,
      unique: true,
      validate: {
        validator: function (email) {
          return /^[^@]+@[^@]+\.[^@]+$/.test(email);
        },
        message: (props) => `${props.value} is not a valid email format !!!.`,
      },
    },
    password: {
      type: String,
      required: true,
      trim: true,
      select: false,
      unique: true,
      minLength: [
        8,
        'Password should be greater than 8 charactor got {VALUE}.',
      ],
      validate: {
        validator: function (password) {
          return /^(?=(?:.*[A-Z]){1,})(?=(?:.*[a-z]){1,})(?=(?:.*\d))(?=(?:.*[!@#$%^&*()]){1,})(?!.(.)\1{2})([A-Za-z0-9!@#$%^&*()\-_=+{};:,<.>]{8,})$/.test(
            password
          );
        },
        message: (props) =>
          `${props.value} is not a valid password format !!!.`,
      },
    },
    confirm_password: {
      type: String,
      required: true,
      trim: true,
      select: false,
      unique: true,
      minLength: [
        8,
        'Password should be greater than 8 charactor got {VALUE}.',
      ],
      validate: {
        validator: function (confirm_password) {
          return confirm_password === this.password;
        },
        message: (props) => `${props.value} is password not match !!!.`,
      },
    },
    contact_number: {
      type: Number,
      required: true,
      select: true,
      unique: true,
    },
    username: {
      type: String,
      required: true,
      trim: true,
      select: true,
    },
    media: {
      type: String,
      trim: true,
      select: true,
    },
    is_verified: {
      type: Boolean,
      required: true,
      select: true,
      default: false,
    },
    product_id: {
      type: mongoose.Types.ObjectId,
      select: true,
      unique: true,
    },
    order_id: {
      type: mongoose.Types.ObjectId,
      select: true,
      unique: true,
    },
    password_reset_token: {
      type: String,
      select: false,
      unique: true,
    },
    password_reset_token_expiry: {
      type: Date,
      select: false,
    },
    deleted_at: {
      type: Date,
      select: false,
      default: null,
    },
  },
  {
    timestamps: true,
    new: true,
    validateBeforeSave: true,
    validateModifiedOnly: true,
  }
);

const User = new mongoose.Model('User', userSchema);

export default User;
