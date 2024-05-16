import PropTypes from 'prop-types';
import { Card, Button } from 'react-bootstrap';
import { useState } from 'react';
import ProductModal from './ProductModal';
import { useContext } from 'react';
import { BasketContext } from '../context/BasketContext';

const ProductCard = ({ product }) => {

    const [showModal, setShowModal] = useState(null);
    const { basket, setBasket } = useContext(BasketContext);

    const handleShowModal = () => {
        setShowModal(true)
    };

    const handleCloseModal = () => {
        setShowModal(false);
    };

    // Adds the product to the basket, and stores in local storage
    // This is bad, because we can edit the stored JSON and change the prices!
    // In real life you'd probably only store a ref to a UUID that you got 
    // from the server, and the basket is stored there, or you'd put it in a cookie with an expiry
    const addItemToBasket = (event) => {
        event.stopPropagation();

        //Check if the product already exists in the basket
        if (basket[product.id]) {
            // If it exists, increase the quantity by 1
            const updatedBasket = {
            ...basket,
            [product.id]: {
                ...basket[product.id],
                quantity: basket[product.id].quantity + 1
            }
            };
            setBasket(updatedBasket);

            //persist to local storage DON'T DO THIS IN A PRODUCTION SYSTEM!
            localStorage.setItem('basket', JSON.stringify(updatedBasket))
        } else {
            // If it doesn't exist, add it to the basket with a quantity of 1
            const updatedBasket = {
            ...basket,
            [product.id]: { ...product, quantity: 1 }
            };
            setBasket(updatedBasket);

            //persist to local storage DON'T DO THIS IN A PRODUCTION SYSTEM!
            localStorage.setItem('basket', JSON.stringify(updatedBasket))
        }
    };

    return (
        <div className="col-md-4 mb-3">
          <Card onClick={handleShowModal}>
            <Card.Img variant="top" src={`https://via.placeholder.com/150?text=${product.name}`} alt={product.name} />
            <Card.Body>
              <Card.Title>{product.name}</Card.Title>
              <Card.Text>Price: £{product.price}</Card.Text>
              <Card.Text>Stock: {product.stock_count}</Card.Text>
              {/* Use onClick handler for the button */}
              <Button variant="primary" onClick={addItemToBasket}>Add to Basket</Button>
            </Card.Body>
            </Card>
            {/* Modal component */}
            <ProductModal
                product={product}
                show={showModal}
                onHide={handleCloseModal}
            />
        </div>
      );
};
  
ProductCard.propTypes = {
    product: 
      PropTypes.shape({
        id: PropTypes.number.isRequired,
        name: PropTypes.string.isRequired,
        price: PropTypes.number.isRequired,
        stock_count: PropTypes.number.isRequired,
      }).isRequired,
  };

export default ProductCard;