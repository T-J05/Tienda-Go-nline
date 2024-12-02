import mongoose from "mongoose";

const userSchema = new mongoose.Schema({
    username: {
        type: String,
        required: true
    },
    contraseña: {
        type: String,
        required: true
    },
    role: {
        type: String,
        required: true,
        default: "admin"
    },
});

const Users = mongoose.model('Users', userSchema);

export default Users;
