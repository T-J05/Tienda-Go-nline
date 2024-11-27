import mongoose from "mongoose";

const productSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true
    },
    price: {
        type: Number,
        required: true
    },
    image: {
        type: String,
        required: true,
        default: "https://www.ip.gov.py/ip/wp-content/uploads/2024/03/2024-03-14-Mondongo.jpg"
    },
});

const Product = mongoose.model('Product', productSchema);

export default Product;
