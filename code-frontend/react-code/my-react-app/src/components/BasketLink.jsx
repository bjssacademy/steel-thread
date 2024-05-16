import { useContext } from 'react';
import { Nav, Badge } from 'react-bootstrap';
import { BasketContext } from '../context/BasketContext';
import { Link } from 'react-router-dom';
import { Basket } from 'react-bootstrap-icons';
import './Basket.css'

const BasketLink = () => {
  const { basket } = useContext(BasketContext);
  
  let totalQuantity = 0;
  for (const item of Object.values(basket)) {
    totalQuantity += item.quantity;
  }
  
    //You can also use reduce if you are familiar with it to get totalQuantity
    //const totalQuantity = Object.values(basket).reduce((total, item) => total + item.quantity, 0);
    
    return (
    <Nav.Link as={Link} to="/basket" className="basket-link">
      <Basket size="32"/>  <Badge variant="primary" className="basketBadge">{totalQuantity}</Badge>
    </Nav.Link>
  );
};

export default BasketLink;