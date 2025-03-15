import { AddressType } from '@/utils';
import mongoose from 'mongoose';

const userAddressSchema = new mongoose.Schema(
  {
    user_id: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'User',
    },
    address_type: {
      type: String,
      enum: [AddressType],
      default: AddressType.PERMANMENT_ADDRESS,
    },
    address_line_1: {
      type: String,
      unique: true,
      select: true,
    },
    address_line_2: {
      type: String,
      unique: true,
      select: true,
    },
    city: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Location',
    },
    state: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Location',
    },
    country: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Location',
    },
    zip_code: {
      type: Number,
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

const UserAddress = new mongoose.Model('UserAddress', userAddressSchema);

export default UserAddress;
